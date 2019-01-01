package main

import (
	"flag"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raahii/arxiv-resources/config"
	"github.com/raahii/arxiv-resources/controller"
	"github.com/raahii/arxiv-resources/db"
)

func setConfig() {
	env := "development"
	flag.Parse()
	if args := flag.Args(); 0 < len(args) && args[0] == "pro" {
		env = "production"
	}
	config.SetEnvironment(env)
}

func main() {
	// read config and open database
	setConfig()
	db.Init()
	defer db.GetConnection().Close()

	// instantiate waf object
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
