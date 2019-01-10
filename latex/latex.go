package latex

import (
	"fmt"
	"strings"
)

func contains(s string, e rune) bool {
	for _, v := range s {
		if e == v {
			return true
		}
	}
	return false
}

func RemoveComment(str string) string {
	str = strings.Replace(str, "%\n", "", -1)

	for {
		if !strings.Contains(str, "%") {
			break
		}

		startIndex := strings.Index(str, "%")
		endIndex := strings.Index(str[startIndex:], "\n")
		if endIndex > 0 {
			str = str[:startIndex] + str[startIndex+endIndex:]
		} else {
			// came end of the document without \n
			return str[:startIndex]
		}
	}

	return str
}

func FindParenthesis(str string, openPare rune) (int, error) {
	pareMap := map[rune]rune{
		'{': '}',
		'[': ']',
	}

	closePare := pareMap[openPare]

	count := 0
	for i, c := range str {
		if c != openPare && c != closePare {
			continue
		}

		// if opening parenthesis
		if c == openPare {
			count++
		}

		// if closing parenthesis
		if c == closePare {
			count--
		}

		if count > 0 {
			continue
		}

		if count == 0 {
			endIndex := i + 1
			return endIndex, nil // success
		} else {
			return -1, fmt.Errorf("Number of parenthesis is mismatch! Closing one is too much!")
		}
	}

	return -1, fmt.Errorf("Number of parenthesis is mismatch! Opening one is too much!")
}

func FindEndOfOneLineCommand(str string, offset int) (int, error) {
	str = str[offset:]

	found := false
	start := -1
	var startChar rune
	for i, c := range str {
		if contains("{[", c) {
			found = true
			start = i
			startChar = c
			break
		}
	}

	if !found {
		return -1, fmt.Errorf("Command not found")
	}
	offset += start
	str = str[start:]

	for {
		end, err := FindParenthesis(str, startChar)
		if err != nil {
			return -1, err
		}

		str = str[end:]
		offset += end

		if len(str) == 0 {
			break
		}

		startChar = []rune(str)[0]
		if startChar != '{' && startChar != '[' {
			break
		}
	}

	return offset, nil
}

func RemoveOneLineCommands(str string, commands []string) (string, error) {
	for _, com := range commands {
		com = com + "{"
		for {
			if !strings.Contains(str, com) {
				break
			}

			startIndex := strings.Index(str, com)
			endIndex, err := FindEndOfOneLineCommand(str, startIndex)
			if err != nil {
				return "", err
			}
			str = str[:startIndex] + str[endIndex:]
		}
	}

	return str, nil
}

func FindMacroCommands(str string) ([]string, error) {
	// read command bellow:
	// \def, \newcommand, \renewcommand,
	commands := []string{
		`\def`,
		`\newcommand`,
		`\renewcommand`,
		`\DeclareMathOperator`,
	}

	// preprocessing
	for _, com := range commands {
		str = strings.Replace(str, com+"*", com, -1)
	}

	followChars := []string{
		`{`,
		`\`,
	}

	macros := []string{}
	for {
		if len(str) == 0 {
			break
		}

		// find position of command start
		startIndex := len(str)
		found := false
		for _, _command := range commands {
			for _, followChar := range followChars {
				if _pos := strings.Index(str, _command+followChar); _pos != -1 {
					if _pos < startIndex {
						found = true
						startIndex = _pos
					}
				}
			}
		}

		if !found {
			// str includes no command anymore
			break
		}

		endIndex, err := FindEndOfOneLineCommand(str, startIndex)
		if err != nil {
			return macros, err
		}
		macro := str[startIndex:endIndex]
		// \newcommand*{...} -> \newcommand{...}
		macros = append(macros, macro)
		str = str[endIndex:]
	}

	return macros, nil
}

func FindEquations(str string) ([]string, error) {
	commands := []string{
		`equation`,
		`align`,
		`aligned`,
		`eqnarray`,
		`subequations`,
	}

	equations := []string{}
	for {
		// find command position
		command := ""
		startIndex := len(str)
		for _, _command := range commands {
			startCommand := fmt.Sprintf(`\begin{%s}`, _command)
			if _pos := strings.Index(str, startCommand); _pos != -1 {
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

		// find equation string
		// ex)
		//     \begin{equation}
		//      [here]
		//     \end{equation}

		// find end of command opening
		endIndex, err := FindEndOfOneLineCommand(str, startIndex)
		if err != nil {
			return []string{""}, err
		}

		str = str[endIndex:]

		// find end of command closing
		endCommand := fmt.Sprintf("\\end{%s", command)
		startIndex = strings.Index(str, endCommand)
		if startIndex == -1 {
			return nil, fmt.Errorf("Corresponding end command is not found!, %s\n%v", command, equations)
		}
		equation := str[:startIndex]

		endIndex, err = FindEndOfOneLineCommand(str, startIndex)
		if err != nil {
			return []string{""}, err
		}
		str = str[endIndex:]

		// remove nested command if exists
		// ex)
		//    \begin{equation}
		//     \begin{align} <- remove
		//       ...
		//     \end{align}   <- remove
		//    \end{equation

		// check nested command exists
		for _, command := range commands {
			startCommand := fmt.Sprintf("\\begin{%s", command)
			if !strings.Contains(equation, startCommand) {
				continue
			}

			// remove command start
			startIndex = strings.Index(equation, startCommand)
			endIndex, err := FindEndOfOneLineCommand(equation, startIndex)
			if err != nil {
				return []string{""}, err
			}
			equation = equation[:startIndex] + equation[endIndex:]

			// remove command end
			endCommand := fmt.Sprintf("\\end{%s", command)
			startIndex = strings.Index(equation, endCommand)
			endIndex, err = FindEndOfOneLineCommand(equation, startIndex)
			if err != nil {
				return []string{""}, err
			}
			equation = equation[:startIndex] + equation[endIndex:]
		}
		equation = strings.Trim(equation, "\n\t ")
		equation = strings.Replace(equation, "\n\n", "\n", -1)
		equations = append(equations, equation)
	}

	return equations, nil
}
