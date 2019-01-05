package config

import (
	yaml "gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
)

var Config Conf

type Environment struct {
	Development Conf `yaml:"development"`
	Production  Conf `yaml:"production"`
}

type Conf struct {
	Database Database `yaml:"db"`
}

type Database struct {
	User     string `yaml:"user"`
	Password string `yaml:"password"`
	Name     string `yaml:"name"`
}

func SetEnvironment(env string) {

	buf, err := ioutil.ReadFile("config/environment.yml")
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
