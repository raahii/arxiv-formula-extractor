package main

import (
	"flag"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raahii/arxiv-equations/config"
	"github.com/raahii/arxiv-equations/controllers"
	"github.com/raahii/arxiv-equations/db"
)

func getPort() string {
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}

	return port
}

func initApp(db *gorm.DB) {
	// Create tables
	models := []interface{}{
		&controllers.Paper{},
		&controllers.Equation{},
		&controllers.Author{},
	}
	for _, model := range models {
		db.AutoMigrate(model)
	}

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
	db.Init()
	database := db.GetConnection()
	initApp(database)
	defer database.Close()

	// instantiate waf object
	e := echo.New()

	// middlewares
	e.Use(middleware.Logger())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{config.Config.Variables["frontAddr"]},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(middleware.Recover())

	// error handler
	e.HTTPErrorHandler = controllers.JSONErrorHandler

	// routing
	e.GET("/papers/:arxiv_id", controllers.FindPaper())

	// start
	err := e.Start(":" + getPort())
	if err != nil {
		log.Fatal(err)
	}
}
