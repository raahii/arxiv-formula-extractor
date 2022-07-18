package database

import (
	"fmt"
	"log"

	"github.com/raahii/arxiv-formula-extractor/config"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB
var err error

func Init() {
	c := config.Config.Database

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		c.User, c.Password, c.Host, c.Name)

	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
}

func GetConnection() *gorm.DB {
	return db
}
