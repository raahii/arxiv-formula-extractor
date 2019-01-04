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

func FindEquations(source string) []string {
	if !strings.Contains(source, "{equation}") {
		return nil
	}

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
