package main

import (
	"flag"
	"log"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/raahii/arxiv-formula-extractor/config"
	"github.com/raahii/arxiv-formula-extractor/controllers"
	"github.com/raahii/arxiv-formula-extractor/database"
	"gorm.io/gorm"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	return port
}

func initApp(DB *gorm.DB) {
	// // Create tables
	// models := []interface{}{
	// 	&controllers.Paper{},
	// 	&controllers.Equation{},
	// 	&controllers.Author{},
	// 	&controllers.Macro{},
	// }
	// for _, model := range models {
	// 	DB.AutoMigrate(model)
	// }

	// Create tarball dirs
	vars := config.Config.Variables
	os.Mkdir(vars["tarballDir"], 0777)
}

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
	database.Init()
	DB := database.GetConnection()
	initApp(DB)

	// instantiate waf object
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.Config.Variables["allowOrigin"]},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())

	// error handler
	e.HTTPErrorHandler = controllers.JSONErrorHandler

	// routing
	e.GET("/papers/:arxiv_id", controllers.FindPaper)

	// start
	err := e.Start(":" + getPort())
	if err != nil {
		log.Fatal(err)
	}
}
