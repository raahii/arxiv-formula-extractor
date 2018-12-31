package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

type Equation struct {
	text string
}

func findEquation(source string) []Equation {
	if !strings.Contains(source, "{equation}") {
		return nil
	}

	r := regexp.MustCompile(`(?s)\\begin\{equation\}(.*?)\\end\{equation\}`)
	m := r.FindAllStringSubmatch(source, -1)

	equations := []Equation{}
	for _, s := range m {
		equations = append(equations, Equation{s[1]})
	}
	return equations
}

func main() {
	data, err := ioutil.ReadFile("samples/tarball/GAN/adversarial.tex")
	if err != nil {
		log.Fatal(err)
	}
	equations := findEquation(string(data))
	for _, s := range equations {
		fmt.Printf("%v\n", s)
	}
}
