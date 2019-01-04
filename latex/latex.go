package latex

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func RemoveComment(str string) string {
	var re = regexp.MustCompile("%.*")
	s := re.ReplaceAllString(str, ``)

	return s
}

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

func FindEquations(source string) []string {
	if !strings.Contains(source, "{equation}") {
		return nil
	}

	pattern := `(?s)\\begin\{(equation|aligned|eqnarray)\}(.*?)\\end\{(equation|aligned|eqnarray)\}`
	re := regexp.MustCompile(pattern)
	m := re.FindAllString(source, -1)

	equations := []string{}
	for _, str := range m {
		str = strings.Replace(str, "{equation}", "{aligned}", -1)
		str = strings.TrimLeft(str, "\n\t")
		str = strings.TrimRight(str, "\n\t")
		str = RemoveTagLines(str, []string{`\label`})
		equations = append(equations, str)
	}
	return equations
}

func main() {
	data, err := ioutil.ReadFile("samples/tarball/GAN/adversarial.tex")
	if err != nil {
		log.Fatal(err)
	}

	latex_source := string(data)
	latex_source = RemoveComment(latex_source)

	equations := FindEquations(latex_source)
	for _, s := range equations {
		fmt.Printf("%v\n", RemoveTagLines(s, []string{`\label`}))
	}
}
