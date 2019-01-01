package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raahii/arxiv-resources/controller"
)

func main() {
	// instantiate echo waf
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// controller
	e.GET("/papers", controller.SearchPapers())
	e.GET("/papers/:id", controller.ShowPaper())

	// start
	e.Start(":1323")
}
