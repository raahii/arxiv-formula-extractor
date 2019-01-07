package latex

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
)

func contains(s []string, e string) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

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

func FindCommandEnd(str string, command string) (int, error) {
	cond := fmt.Sprintf("\nstr:\n'%s'\n command:\n'%s'", str, command)
	if str[:len(command)] != command {
		return -1, fmt.Errorf("Passed str is not started command!" + cond)
	}

	ignore := false
	nBraces := 0
	for i, c := range str {
		if c != rune('{') && c != rune('}') {
			continue
		}

		if c == rune('{') {
			if i == len(command) {
				// first brace
				ignore = true
			} else {
				nBraces++
			}
			continue
		}

		if c == rune('}') {
			if ignore {
				ignore = false
				continue
			}

			nBraces--
			if nBraces > 0 {
				continue
			} else if nBraces == 0 {
				return i + 1, nil // success
			} else {
				return -1, fmt.Errorf("Number of braces is mismatch! } is too much!" + cond)
			}
		}
	}

	if nBraces > 0 {
		return -1, fmt.Errorf("Number of braces is mismatch! { is too much!" + cond)
	} else {
		return -1, fmt.Errorf("Brace is not found" + cond)
	}

}

func FindMacros(str string) ([]string, error) {
	// read command bellow:
	// \def, \newcommand, \renewcommand,
	commands := []string{
		`\def`,
		`\newcommand`,
		`\renewcommand`,
	}
	followChars := []string{
		`{`,
		`*`,
		`\`,
	}

	macros := []string{}
	for {
		if len(str) == 0 {
			break
		}

		command := ""
		startIndex := len(str)
		for _, _command := range commands {
			for _, followChar := range followChars {
				if _pos := strings.Index(str, _command+followChar); _pos != -1 {
					if _pos < startIndex {
						startIndex = _pos
						command = _command
					}
				}
			}
		}

		if command == "" {
			// str includes no command anymore
			break
		}

		endIndex, err := FindCommandEnd(str[startIndex:], command)
		if err != nil {
			return macros, err
		}
		log.Println(str[startIndex : startIndex+endIndex])
		macros = append(macros, str[startIndex:startIndex+endIndex])
		str = str[startIndex+endIndex:]
	}

	return macros, nil
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
