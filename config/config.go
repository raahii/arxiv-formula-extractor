package config

import (
	"io/ioutil"
	"log"
	"os"

	yaml "gopkg.in/yaml.v2"
)

var Config Conf

type Conf struct {
	Environment string
	Database    Database
	Variables   map[string]string `yaml:"vars"`
}

type Database struct {
	User     string
	Password string
	Name     string
	Host     string
}

func SetEnvironment(env string) {
	Config = Conf{}

	// env
	env = os.Getenv("ENV")
	if env == "" {
		Config.Environment = "development"
	} else {
		Config.Environment = env
	}

	// database setting
	db := Database{}
	db.User = os.Getenv("DB_USER")
	db.Password = os.Getenv("DB_PASS")
	db.Name = os.Getenv("DB_NAME")
	db.Host = os.Getenv("DB_HOST")
	Config.Database = db

	// other variables
	buf, err := ioutil.ReadFile("config/" + Config.Environment + ".yml")
	if err != nil {
		log.Fatal(err)
	}

	err = yaml.Unmarshal(buf, &Config.Variables)
	if err != nil {
		log.Fatal(err)
	}
}
