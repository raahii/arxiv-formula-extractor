package latex

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func RemoveComment(str string) string {
	newStr := make([]rune, 0, len([]rune(str)))
	percent := rune('%')
	newline := rune('\n')

	valid := true
	for _, c := range str {
		if c == percent {
			valid = false
		}

		if c == newline {
			valid = true
		}

		if valid {
			newStr = append(newStr, c)
		}
	}

	return string(newStr)
}

func RemoveTags(str string, tags []string) string {
	tagsStr := strings.Join(tags, "|")
	pattern := fmt.Sprintf(`\\(%s)(\{.*\})?(\n)?`, tagsStr)
	re := regexp.MustCompile(pattern)

	return re.ReplaceAllString(str, "")
}

func FindMacros(source string) []string {
	// read tags bellow:
	// \def, \newcommand, \renewcommand,
	// \newenvironment, \renewenvironment
	tags := []string{
		`\def`,
		`\newcommand`,
		`\renewcommand`,
		`\newenvironment`,
		`\renewenvironment`,
	}

	pattern := fmt.Sprintf(`(\%s)`, strings.Join(tags, `|\`)) + `\*?{.*`
	re := regexp.MustCompile(pattern)
	matchStrs := re.FindAllString(source, -1)

	for i := 0; i < len(matchStrs); i++ {
		for _, tag := range tags {
			if strings.Contains(matchStrs[i], tag+"*{") {
				matchStrs[i] = strings.Replace(matchStrs[i], tag+"*{", tag+"{", 1)
				break
			}
		}
	}

	return matchStrs
}

func FindEquations(source string) []string {
	if !strings.Contains(source, "{equation}") {
		return nil
	}

	// TODO: change not to use regular expressions
	pattern := `(?s)\\begin\{(equation|align|aligned|eqnarray)\}(.*?)\\end\{(equation|align|aligned|eqnarray)\}`
	re := regexp.MustCompile(pattern)
	m := re.FindAllStringSubmatch(source, -1)

	equations := []string{}
	for _, strs := range m {
		eq := strs[2]
		eq = strings.TrimLeft(eq, "\n\t")
		eq = strings.TrimRight(eq, "\n\t")
		eq = RemoveTags(eq, []string{"label", "nonumber"})
		equations = append(equations, eq)
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
		fmt.Printf("%v\n", RemoveTags(s, []string{`\label`}))
	}
}
