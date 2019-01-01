package main

import (
	"flag"
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/raahii/arxiv-resources/config"
	"github.com/raahii/arxiv-resources/controller"
	"github.com/raahii/arxiv-resources/db"
)

func InitTables(db *gorm.DB) {

	// Create tables
	if !db.HasTable(&controller.Paper{}) {
		db.CreateTable(&controller.Paper{})
	}
	paper := controller.Paper{}
	paper.Id = 1
	paper.Title = "Sample Paper"
	paper.Authors = "Mr.hogehoge"
	paper.Version = 1
	paper.Abstract = "Abstract here."
	paper.Equations = []controller.Equation{
		controller.Equation{1, "y=ax"},
		controller.Equation{2, "F=ma"},
	}
	db.Create(&paper)
	if !db.HasTable(&controller.Equation{}) {
		db.CreateTable(&controller.Equation{})
	}

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
	InitTables(database)
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
	e.GET("/papers", controller.SearchPapers())
	e.GET("/papers/:id", controller.ShowPaper())

	// start
	e.Start(":1323")
}
