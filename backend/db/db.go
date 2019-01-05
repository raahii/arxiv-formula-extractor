package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/raahii/arxiv-equations/backend/config"
	"log"
)

var db *gorm.DB
var err error

func Init() {
	c := config.Config.Database

	user := c.User
	password := c.Password
	dbname := c.Name

	db, err = gorm.Open("mysql", user+":"+password+"@/"+dbname+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}
