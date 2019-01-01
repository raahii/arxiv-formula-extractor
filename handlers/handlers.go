package handlers

import (
	"github.com/labstack/echo"
	arxiv "github.com/raahii/latexeq-copier/arxiv"
	"net/http"
)

func TestPage() echo.HandlerFunc {
	return func(c echo.Context) error {
		return c.JSON(http.StatusOK, "Hello World")
	}
}

func SearchPaper() echo.HandlerFunc {
	return func(c echo.Context) error {
		papers := []arxiv.Entry{arxiv.Entry{}}
		return c.JSON(http.StatusOK, papers)
	}
}
