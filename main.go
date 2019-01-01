package main

import (
	"./handlers"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
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

	// handlers
	e.GET("/hello", handlers.TestPage())
	e.GET("/papers", handlers.SearchPapers())
	e.GET("/papers/:id", handlers.ShowPaper())

	// start
	e.Start(":1323")
}
