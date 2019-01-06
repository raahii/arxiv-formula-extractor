package main

import (
	"flag"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raahii/arxiv-equations/backend/config"
	"github.com/raahii/arxiv-equations/backend/controller"
	"github.com/raahii/arxiv-equations/backend/db"
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
	os.Mkdir("tarballs", 0777)
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
	e.GET("/papers", controller.FindPaperFromUrl())
	e.GET("/papers/:id", controller.ShowPaper())

	// start
	e.Start(":1323")
}
