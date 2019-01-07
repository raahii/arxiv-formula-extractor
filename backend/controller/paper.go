package controller

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"strconv"
	"strings"

	"github.com/labstack/echo"
	zglob "github.com/mattn/go-zglob"
	"github.com/raahii/arxiv-equations/backend/arxiv"
	"github.com/raahii/arxiv-equations/backend/db"
	"github.com/raahii/arxiv-equations/backend/latex"
)

func readFile(path string) string {
	str, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	return string(str)
}

func readAllSources(mainLatexPath string, basePath string) string {
	// read all \input or \include tag and
	// obtain all related sources concatenated string
	source := readFile(mainLatexPath)
	source = latex.RemoveComment(source)
	source = strings.Replace(source, "*{", "{", -1)
	source = strings.Replace(source, "*}", "}", -1)

	re := regexp.MustCompile(`\\(input|include)\{(.*?)\}`)

	resolveInputTag := func(s string) string {
		path := re.FindStringSubmatch(s)[2]
		if filepath.Ext(path) == "" {
			path = path + ".tex"
		}
		_source := readFile(filepath.Join(basePath, path))
		_source = latex.RemoveComment(_source)
		return _source
	}

	// # TODO: improve efficiency
	for {
		if re.FindAllString(source, 1) == nil {
			break
		}
		source = re.ReplaceAllStringFunc(source, resolveInputTag)
	}

	return source
}

func findSourceRoot(paths []string) string {
	// search source which includes '\documentclass'
	found := false
	mainPath := ""
	for _, path := range paths {
		source := readFile(path)
		source = latex.RemoveComment(source)
		if strings.Contains(source, `\documentclass`) {
			found = true
			mainPath = path
		}
	}
	if !found {
		log.Fatal(fmt.Errorf("Main latex source is not found"))
	}
	return mainPath
}

func extractArxivId(arxivUrl string) string {
	// ex) https://arxiv.org/abs/1406.2661
	strs := strings.Split(arxivUrl, "/")
	return strs[len(strs)-1]
}

func (paper *Paper) readLatexSource(path string) {
	var err error

	// download tarball
	log.Println("Downloading tarball", paper.TarballUrl)
	tarballPath := filepath.Join(path, paper.ArxivId+".tar.gz")
	err = arxiv.DownloadTarball(paper.TarballUrl, tarballPath)
	if err != nil {
		log.Fatal(err)
	}

	// decompress tarball
	sourcePath := filepath.Join(path, paper.ArxivId)
	log.Println("Decompressing tarball", sourcePath)
	os.Mkdir(sourcePath, 0777)

	err = exec.Command("tar", "-xvzf", tarballPath, "-C", sourcePath).Run()
	if err != nil {
		log.Fatal(err)
	}

	// list all *.tex
	log.Println("Processing tex files")
	pattern := filepath.Join(sourcePath, "**/*.tex")
	files, err := zglob.Glob(pattern)
	if err != nil {
		log.Fatal(err)
	}

	// find root latex source file
	rootFile := findSourceRoot(files)

	// obtain all latex source
	allSource := readAllSources(rootFile, sourcePath)

	// obtain macros
	log.Println("Extracting macros")
	macros, err := latex.FindMacros(allSource)
	if err != nil {
		log.Fatal(err)
	}
	paper.Macros = strings.Join(macros, "\n")

	// obtain equations
	log.Println("Extracting equations")
	equationStrs, err := latex.FindEquations(allSource)
	if err != nil {
		log.Fatal(err)
	}
	equations := []Equation{}
	for _, str := range equationStrs {
		eq := Equation{}
		eq.Expression = str
		equations = append(equations, eq)
	}
	paper.Equations = equations
}

func FetchPaper(arxivId string) Paper {
	// search papers
	params := map[string]string{
		"id_list": arxivId,
	}
	apiResult := arxiv.SearchPapers(params)
	apiEntry := apiResult.Entries[0]

	// convert api result to paper entity
	authors := []Author{} // for now, authors are just a string
	for _, a := range apiEntry.Authors {
		author := Author{}
		author.Name = a.Name
		authors = append(authors, author)
	}

	// extract urls
	absUrl, tarballUrl := "", ""
	for _, link := range apiEntry.Links {
		if link.Type == "text/html" {
			absUrl = link.Value
			tarballUrl = strings.Replace(absUrl, "/abs", "/e-print", 1)
			break
		}
	}

	// make a paper entitiy
	paper := Paper{}
	paper.ArxivId = arxivId
	paper.Authors = authors
	paper.Title = apiEntry.Title
	paper.Abstract = apiEntry.Summary
	paper.AbsUrl = absUrl
	paper.TarballUrl = tarballUrl

	return paper
}

func FindPaperFromUrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		// obtain url from GET parameters
		url := c.QueryParam("url")
		if url == "" {
			log.Fatal(fmt.Errorf("Invalid parameters"))
		}

		// remove version number from url
		r := regexp.MustCompile(`v[1-9]+$`)
		url = r.ReplaceAllString(url, "")

		// convert paper url to id on arxiv, id on this app.
		arxivId := extractArxivId(url)

		// find the paper
		database := db.GetConnection()
		paper := Paper{}
		if database.Where("arxiv_id = ?", arxivId).First(&paper).RecordNotFound() {
			// if the paper doesn't exist in the database

			// fetch the paper
			paper = FetchPaper(arxivId)

			// extract macros and equations
			tarballDir := "tarballs"
			paper.readLatexSource(tarballDir)

			if dbc := database.Create(&paper); dbc.Error != nil {
				log.Fatal(dbc.Error)
			}
		} else {
			database.Model(&paper).Related(&paper.Equations).Related(&paper.Authors)
			// tarballDir := "tarballs"
			// paper.readLatexSource(tarballDir)
		}

		// add macro to process fine for unsupported command in mathjax
		defaultMacros := []string{
			`\newcommand{\bm}[1]{\boldsymbol #1}`,
		}
		paper.Macros += "\n" + strings.Join(defaultMacros, "\n")

		response := map[string]interface{}{
			"paper": paper,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func ShowPaper() echo.HandlerFunc {
	return func(c echo.Context) error {
		// obtain path parameter
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}

		// find the paper
		db := db.GetConnection()
		paper := Paper{}
		db.First(&paper, int32(id))

		response := map[string]interface{}{
			"paper": paper,
		}

		return c.JSON(http.StatusOK, response)
	}
}
