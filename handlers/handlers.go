package handlers

import (
	"github.com/labstack/echo"
	arxiv "github.com/raahii/latexeq-copier/arxiv"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func TestPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	}
}

func toIntIdAndVersion(apiId string) (int32, int32) {
	// ex) http://arxiv.org/abs/cond-mat/9511068v1
	urlElems := strings.Split(apiId, "/")
	lastElem := urlElems[len(urlElems)-1]
	paperStr := strings.Split(lastElem, "v")
	id, err := strconv.Atoi(strings.Replace(paperStr[0], ".", "", -1))
	if err != nil {
		log.Fatal(err)
	}
	ver, err := strconv.Atoi(paperStr[1])
	if err != nil {
		log.Fatal(err)
	}

	return int32(id), int32(ver)
}

func SearchPapers() echo.HandlerFunc {
	return func(c echo.Context) error {
		query := c.QueryParam("q")
		apiResult := arxiv.SearchPapers(query)
		apiEntries := apiResult.Entries

		// convert api result to paper entity
		papers := []Paper{}
		for _, entry := range apiEntries {
			id, ver := toIntIdAndVersion(entry.Id)

			authors := []string{} // for now, authors are just a string
			for _, a := range entry.Authors {
				authors = append(authors, a.Name)
			}

			// make a paper entitiy
			paper := Paper{}
			paper.Id = id
			paper.Version = ver
			paper.Authors = strings.Join(authors, " ")
			paper.Title = entry.Title
			paper.Abstract = entry.Summary

			papers = append(papers, paper)
		}

		response := map[string]interface{}{
			"query":     query,
			"n_results": len(papers),
			"papers":    papers,
		}

		return c.JSON(http.StatusOK, response)
	}
}

func ShowPaper() echo.HandlerFunc {
	return func(c echo.Context) error {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			log.Fatal(err)
		}
		paper := Paper{}
		paper.Id = int32(id)
		paper.Title = "sample paper"

		response := map[string]interface{}{
			"paper": paper,
		}

		return c.JSON(http.StatusOK, response)
	}
}
