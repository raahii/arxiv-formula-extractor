package latex

import (
	"fmt"
	"regexp"
	"strings"
)

func RemoveComment(str string) string {
	for {
		if !strings.Contains(str, "%") {
			break
		}

		startIndex := strings.Index(str, "%")
		endIndex := startIndex + strings.Index(str[startIndex:], "\n")
		str = str[:startIndex] + str[endIndex:]
	}

	return str
}

func RemoveTags(str string, tags []string) (string, error) {
	tagsStr := strings.Join(tags, "|")
	pattern := fmt.Sprintf(`\\(%s)(\{.*\})?(\n)?`, tagsStr)
	re, err := regexp.Compile(pattern)
	if err != nil {
		return "", err
	}

	return re.ReplaceAllString(str, ""), nil
}

func FindMacroCommandEnd(str string, command string) (int, error) {
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
		`\DeclareMathOperator`,
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

		endIndex, err := FindMacroCommandEnd(str[startIndex:], command)
		if err != nil {
			return macros, err
		}
		macros = append(macros, str[startIndex:startIndex+endIndex])
		str = str[startIndex+endIndex:]
	}

	return macros, nil
}

func FindEquations(str string) ([]string, error) {
	commands := []string{
		`equation`,
		`align`,
		`aligned`,
		`eqnarray`,
		`array`,
	}

	equations := []string{}
	for {
		command := ""
		startIndex := len(str)
		for _, _command := range commands {
			commandStart := fmt.Sprintf("\\begin{%s", _command)
			if _pos := strings.Index(str, commandStart); _pos != -1 {
				if _pos < startIndex {
					startIndex = _pos
					command = _command
				}
			}
		}

		if command == "" {
			// str includes no command anymore
			break
		}

		// find end of command opening
		flag := false
		for i, c := range str[startIndex:] {
			if flag {
				startIndex += i
				break
			}
			if c == rune('\n') {
				flag = true
			}
		}
		str = str[startIndex:]

		// find end of command closing
		commandEnd := fmt.Sprintf("\\end{%s", command)
		endIndex := strings.Index(str, commandEnd)
		if endIndex == -1 {
			return nil, fmt.Errorf("Corresponding end command is not found!, %s\n%v", command, equations)
		}

		// check nested command exists
		equation := str[:endIndex]
		var s, e int
		for _, command := range commands {
			if !strings.Contains(equation, command) {
				continue
			}
			// remove command start
			startCommand := fmt.Sprintf("\\begin{%s", command)
			s = strings.Index(equation, startCommand)
			e = s + strings.Index(equation[s:], "\n")
			equation = equation[:s] + equation[e:]

			// remove command end
			endCommand := fmt.Sprintf("\\end{%s", command)
			s = strings.Index(equation, endCommand)
			e = s + strings.Index(equation[s:], "\n")
			if e == -1 {
				equation = equation[:s]
			} else {
				equation = equation[:s] + equation[e:]
			}
		}
		equation = strings.Trim(equation, "\n\t ")
		equations = append(equations, equation)
		str = str[endIndex:]
	}

	return equations, nil
}
