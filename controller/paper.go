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
	"strings"

	"github.com/labstack/echo"
	zglob "github.com/mattn/go-zglob"
	"github.com/raahii/arxiv-equations/arxiv"
	"github.com/raahii/arxiv-equations/config"
	"github.com/raahii/arxiv-equations/db"
	"github.com/raahii/arxiv-equations/latex"
)

func readFile(path string) (string, error) {
	str, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func readAllSources(mainLatexPath string, basePath string) (string, error) {
	// read all \input or \include tag and
	// obtain all related sources concatenated string
	source, err := readFile(mainLatexPath)
	if err != nil {
		return "", err
	}

	source = latex.RemoveComment(source)
	source = strings.Replace(source, "*{", "{", -1)
	source = strings.Replace(source, "*}", "}", -1)

	re, err := regexp.Compile(`\\(input|include)\{(.*?)\}`)
	if err != nil {
		return "", err
	}

	resolveInputTag := func(s string) string {
		path := re.FindStringSubmatch(s)[2]
		if filepath.Ext(path) == "" {
			path = path + ".tex"
		}
		_source, err := readFile(filepath.Join(basePath, path))
		if err != nil {
			panic(err)
		}
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

	return source, nil
}

func findSourceRoot(paths []string) (string, error) {
	// search source which includes '\documentclass'
	found := false
	mainPath := ""
	for _, path := range paths {
		source, err := readFile(path)
		if err != nil {
			return "", err
		}
		source = latex.RemoveComment(source)
		if strings.Contains(source, `\documentclass`) {
			found = true
			mainPath = path
		}
	}
	if !found {
		return "", fmt.Errorf("Latex file is not found")
	}
	return mainPath, nil
}

func extractArxivId(arxivUrl string) string {
	// ex) https://arxiv.org/abs/1406.2661
	strs := strings.Split(arxivUrl, "/")
	return strs[len(strs)-1]
}

func (paper *Paper) readLatexSource(path string) error {
	var err error

	// download tarball
	tarballPath := filepath.Join(path, paper.ArxivId+".tar.gz")
	err = arxiv.DownloadTarball(paper.TarballUrl, tarballPath)
	if err != nil {
		return newErrorWithMsg(err, "Error occured during downloading tarball")
	}

	// decompress tarball
	sourcePath := filepath.Join(path, paper.ArxivId)
	os.Mkdir(sourcePath, 0777)

	err = exec.Command("tar", "-xvzf", tarballPath, "-C", sourcePath).Run()
	if err != nil {
		return newErrorWithMsg(err, "Error occured during decompressing tarball.")
	}

	// list all *.tex
	pattern := filepath.Join(sourcePath, "**/*.tex")
	files, err := zglob.Glob(pattern)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during processing tex files(1)")
	}

	// find root latex source file
	rootFile, err := findSourceRoot(files)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during processing tex files(2)")
	}

	// obtain all latex source
	allSource, err := readAllSources(rootFile, sourcePath)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during processing tex files(3)")
	}

	// obtain macros
	log.Println("Extracting macros")
	macros, err := latex.FindMacros(allSource)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during extracting macros")
	}
	paper.Macros = strings.Join(macros, "\n")

	// obtain equations
	log.Println("Extracting equations")
	equationStrs, err := latex.FindEquations(allSource)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during extracting equations")
	}
	equations := []Equation{}
	for _, str := range equationStrs {
		eq := Equation{}
		eq.Expression = str
		equations = append(equations, eq)
	}
	paper.Equations = equations

	// remove tarball
	if err := os.Remove(tarballPath); err != nil {
		return err
	}
	if err := os.RemoveAll(sourcePath); err != nil {
		return err
	}

	return nil
}

func FetchPaper(arxivId string) (Paper, error) {
	// search paper from id
	params := map[string]string{
		"id_list": arxivId,
	}
	apiResult, err := arxiv.SearchPapers(params)
	if err != nil {
		return Paper{}, err
	}

	apiEntry := apiResult.Entries[0]
	if apiEntry.Title == "Error" {
		err := fmt.Errorf(apiEntry.Summary)
		return Paper{}, err
	}

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

	return paper, nil
}

func FindPaperFromUrl() echo.HandlerFunc {
	return func(c echo.Context) error {
		// obtain url from GET parameters
		url := c.QueryParam("url")
		if url == "" {
			fmt.Errorf("Invalid parameters")
		}

		// remove version number from url
		r := regexp.MustCompile(`v[1-9]+$`)
		url = r.ReplaceAllString(url, "")

		// convert paper url to id on arxiv, id on this app.
		arxivId := extractArxivId(url)

		// find the paper
		paper := Paper{}
		database := db.GetConnection()
		if database.Where("arxiv_id = ?", arxivId).First(&paper).RecordNotFound() {
			// if the paper doesn't exist in the database, fetch the paper
			_paper, err := FetchPaper(arxivId)
			if err != nil {
				return err
			}

			// extract macros and equations
			vars := config.Config.Variables
			tarballDir := vars["tarballDir"]
			_paper.readLatexSource(tarballDir)

			if dbc := database.Create(&_paper); dbc.Error != nil {
				return dbc.Error
			}
			paper = _paper
		} else {
			database.Model(&paper).Related(&paper.Equations).Related(&paper.Authors)
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
