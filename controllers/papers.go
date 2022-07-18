package controllers

import (
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/mattn/go-zglob"
	"github.com/raahii/arxiv-formula-extractor/arxiv"
	"github.com/raahii/arxiv-formula-extractor/config"
	"github.com/raahii/arxiv-formula-extractor/database"
	"github.com/raahii/arxiv-formula-extractor/latex"
	"gorm.io/gorm"
)

func FindPaper(c echo.Context) error {
	DB := database.GetConnection()

	// obtain url from GET parameters
	arxivId := c.Param("arxiv_id")

	// find the paper
	paper := Paper{}
	err := DB.Where("arxiv_id = ?", arxivId).First(&paper).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		// if the paper doesn't exist in the database, fetch the paper
		paper, err := FetchPaper(arxivId)
		if err != nil {
			return err
		}

		// extract macros and equations
		vars := config.Config.Variables
		tarballDir := vars["tarballDir"]
		err = paper.readLatexSource(tarballDir)
		if err != nil {
			return err
		}

		if dbc := DB.Create(&paper); dbc.Error != nil {
			return dbc.Error
		}
	} else {
		DB.Where("paper_id = ?", paper.ID).Find(&paper.Equations)
		DB.Where("paper_id = ?", paper.ID).Find(&paper.Authors)
		DB.Where("paper_id = ?", paper.ID).Find(&paper.Macros)
	}

	response := map[string]interface{}{
		"paper": paper,
	}

	return c.JSON(http.StatusOK, response)
}

func readFile(path string) (string, error) {
	str, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	return string(str), nil
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findSourceRoot(paths []string) ([]string, error) {
	// search source which includes '\documentclass'
	candidates := []string{}
	for _, path := range paths {
		source, err := readFile(path)
		if err != nil {
			return []string{}, err
		}
		source = latex.RemoveComment(source)
		if strings.Contains(source, `\documentclass`) {
			candidates = append(candidates, path)
		}
	}
	if len(candidates) > 0 {
		return candidates, nil
	} else {
		return []string{}, fmt.Errorf("Root latex file is not found")
	}
}

func resolveInputs(mainLatexPath string, basePath string) (string, error) {
	// read all \input or \include tag and
	// obtain all related sources concatenated string
	source, err := readFile(mainLatexPath)
	if err != nil {
		return "", err
	}

	source = latex.RemoveComment(source)
	source = strings.Replace(source, "*{", "{", -1)
	source = strings.Replace(source, "*}", "}", -1)

	commands := []string{
		`\input`,
		`\include`,
		`\usepackage`,
	}

	// replace the command with actual file content
	for _, com := range commands {
		com = com + "{"
		for {
			if !strings.Contains(source, com) {
				break
			}

			// find command
			startIndex := strings.Index(source, com)
			endIndex, err := latex.FindEndOfOneLineCommand(source, startIndex)
			if err != nil {
				return "", err
			}

			// read path in the brace
			path := source[startIndex+len(com) : endIndex-1]
			if filepath.Ext(path) == "" {
				if strings.Contains(com, `\usepackage`) {
					path = path + ".sty"
				} else {
					path = path + ".tex"
				}
			}

			path = filepath.Join(basePath, path)

			if _, err := os.Stat(path); os.IsNotExist(err) {
				// remove the command
				source = source[:startIndex] + source[endIndex:]
				continue
			}

			// read file content
			_source, err := readFile(path)
			if err != nil {
				return "", err
			}

			// replace
			source = source[:startIndex] + _source + source[endIndex:]
		}
	}

	return source, nil
}

func readAllSources(latexFiles []string, basePath string) (string, error) {
	// find candidates for the root latex file
	rootFiles, err := findSourceRoot(latexFiles)
	if err != nil {
		return "", err
	}

	// resolve \input, \include commands for each root file
	allSources := []string{}
	for _, rootFile := range rootFiles {
		source, err := resolveInputs(rootFile, basePath)
		if err != nil {
			return "", err
		}
		allSources = append(allSources, source)
	}

	// if one candidate found, return the source
	if len(allSources) == 1 {
		return allSources[0], nil
	}

	// if multiple candiates found, the most longest source is
	// thought to be main latex file...
	longestSource := ""
	for _, source := range allSources {
		if len(source) > len(longestSource) {
			longestSource = source
		}
	}
	return longestSource, nil
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

	// obtain all latex source
	allSource, err := readAllSources(files, sourcePath)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during processing tex files(3)")
	}

	// remove comment and \label command
	allSource = latex.RemoveComment(allSource)
	allSource, err = latex.RemoveOneLineCommands(allSource, []string{`\label`})
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during removing unnecessary commands")
	}

	// obtain macros
	macroStrs, err := latex.FindMacroCommands(allSource)
	if err != nil {
		return newErrorWithMsg(err, "Error occurred during extracting macros")
	}
	macros := []Macro{}
	for _, str := range macroStrs {
		if strings.Contains(str, "@") {
			continue
		}
		// extract command part, ex) \def\bm... -> \bm
		s := 3 + strings.Index(str[3:], `\`)
		e := len(str)
		for _, c := range []string{` `, `}`, `{`, `[`} {
			if strings.Contains(str, c) {
				e = min(s+strings.Index(str[s:], c), e)
			}
		}
		command := str[s:e]

		macro := Macro{}
		macro.Command = command
		macro.Expression = str
		macros = append(macros, macro)
	}
	paper.Macros = macros

	// obtain equations
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
	os.Remove(tarballPath)
	os.RemoveAll(sourcePath)

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
