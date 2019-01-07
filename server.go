package main

import (
	"flag"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raahii/arxiv-equations/config"
	"github.com/raahii/arxiv-equations/controller"
	"github.com/raahii/arxiv-equations/db"
)

func initApp(db *gorm.DB) {
	// Create tables
	models := []interface{}{
		&controller.Paper{},
		&controller.Equation{},
		&controller.Author{},
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
	e.Use(middleware.Recover())
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	// controller
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/papers", controller.FindPaperFromUrl())

	// start
	port := os.Getenv("PORT")
	if port == "" {
		port = "1323"
	}
	e.Start(":" + port)
}
