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

	// handlers
	e.GET("/hello", handlers.TestPage())
	e.GET("/papers", handlers.SearchPaper())

	// start
	e.Start(":1323")
}
