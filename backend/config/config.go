package config

import (
	"io/ioutil"
	"log"

	yaml "gopkg.in/yaml.v2"
)

var Config Conf

type Environment struct {
	Development Conf `yaml:"development"`
	Production  Conf `yaml:"production"`
}

type Conf struct {
	Database  Database          `yaml:"db"`
	Variables map[string]string `yaml:"vars"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
	Host     string `yaml:"host"`
}

func SetEnvironment(env string) {

	buf, err := ioutil.ReadFile("config/config.yml")
	if err != nil {
		log.Fatal(err)
	}

	var environment Environment

	err = yaml.Unmarshal(buf, &environment)
	if err != nil {
		log.Fatal(err)
	}

	if env == "development" {
		Config = environment.Development
	} else {
		Config = environment.Production
	}
}
