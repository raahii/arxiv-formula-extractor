package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func RemoveTagLines(str string, tags []string) string {
	lines := strings.Split(str, "\n")

	var found bool
	new_lines := []string{}
	for _, line := range lines {
		found = false
		for _, tag := range tags {
			if strings.Contains(line, tag) {
				found = true
				break
			}
		}
		if !found {
			new_lines = append(new_lines, line)
		}
	}

	return strings.Join(new_lines, "\n")
}

func FindEquations(source string) []Equation {
	if !strings.Contains(source, "{equation}") {
		return nil
	}

	r := regexp.MustCompile(`(?s)\\begin\{equation\}(.*?)\\end\{equation\}`)
	m := r.FindAllStringSubmatch(source, -1)

	equations := []Equation{}
	for _, s := range m {
		str := strings.Trim(s[1], "\n\t")
		equations = append(equations, Equation{str})
	}
	return equations
}

func main() {
	data, err := ioutil.ReadFile("samples/tarball/GAN/adversarial.tex")
	if err != nil {
		log.Fatal(err)
	}

	equations := FindEquations(string(data))
	for _, s := range equations {
		fmt.Printf("%v\n", s)
	}
}
