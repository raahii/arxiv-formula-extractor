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
	fmt.Println(pattern)
	re := regexp.MustCompile(pattern)
	matchStrs := re.FindAllString(source, -1)

	for i := 0; i < len(matchStrs); i++ {
		fmt.Printf(string(i))
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
	m := re.FindAllString(source, -1)

	equations := []string{}
	re = regexp.MustCompile("(equation|align|eqnarray)")
	for _, str := range m {
		str = re.ReplaceAllString(str, "aligned")
		str = strings.TrimLeft(str, "\n\t")
		str = strings.TrimRight(str, "\n\t")
		str = RemoveTags(str, []string{"label", "nonumber"})
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
		fmt.Printf("%v\n", RemoveTags(s, []string{`\label`}))
	}
}
