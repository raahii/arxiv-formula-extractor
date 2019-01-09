package latex

import (
	"fmt"
	"strings"
)

func FindCommandEnd(str string) (int, error) {
	// ignore first open brace
	found := false
	start := 0
	for i, c := range str {
		if c == '{' {
			found = true
			start = i + 1
			break
		}
	}

	if !found {
		return -1, fmt.Errorf("String doesn't have any command")
	}

	count := 1 // find <count> closing braces
	for i, c := range str[start:] {
		if c != rune('{') && c != rune('}') {
			continue
		}

		if c == rune('{') {
			count++
			continue
		}

		if c == rune('}') {
			count--
		}

		if count > 0 {
			continue
		}

		if count == 0 {
			return start + i + 1, nil // success
		} else {
			return -1, fmt.Errorf("Number of braces is mismatch! '}' is too much!")
		}
	}

	return -1, fmt.Errorf("Number of braces is mismatch! '{' is too much!")
}

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

func RemoveSimpleCommands(str string, commands []string) (string, error) {
	for _, com := range commands {
		com = com + "{"
		for {
			if !strings.Contains(str, com) {
				break
			}

			startIndex := strings.Index(str, com)
			endIndex, err := FindCommandEnd(str[startIndex:])
			if err != nil {
				return "", err
			}
			endIndex += startIndex
			str = str[:startIndex] + str[endIndex:]
		}
	}

	return str, nil
}

func FindMacroCommandEnd(str string, command string) (int, error) {
	if str[:len(command)] != command {
		return -1, fmt.Errorf("Passed str is not started command!")
	}

	// for \newcommand, \renewcommand
	endIndex := 0
	if str[len(command)] == '{' {
		firstBraceEnd, err := FindCommandEnd(str)
		if err != nil {
			return -1, err
		}
		str = str[firstBraceEnd:]
		endIndex += firstBraceEnd
	}

	end, err := FindCommandEnd(str)
	if err != nil {
		return -1, err
	}

	endIndex += end
	return endIndex, nil
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
		macro := str[startIndex : startIndex+endIndex]
		macro = strings.Replace(macro, command+"*", command, 1)
		macros = append(macros, macro)
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
		endIndex, err := FindCommandEnd(str[startIndex:])
		if err != nil {
			return []string{""}, err
		}
		endIndex += startIndex
		str = str[endIndex:]

		// find end of command closing
		endCommand := fmt.Sprintf("\\end{%s", command)
		startIndex = strings.Index(str, endCommand)
		if startIndex == -1 {
			return nil, fmt.Errorf("Corresponding end command is not found!, %s\n%v", command, equations)
		}
		equation := str[:startIndex]

		endIndex, err = FindCommandEnd(str[startIndex:])
		if err != nil {
			return []string{""}, err
		}
		endIndex += startIndex
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
			endIndex, err := FindCommandEnd(equation[startIndex:])
			if err != nil {
				return []string{""}, err
			}
			endIndex += startIndex
			equation = equation[:startIndex] + equation[endIndex:]

			// remove command end
			endCommand := fmt.Sprintf("\\end{%s", command)
			startIndex = strings.Index(equation, endCommand)
			endIndex, err = FindCommandEnd(equation[startIndex:])
			if err != nil {
				return []string{""}, err
			}
			endIndex += startIndex
			equation = equation[:startIndex] + equation[endIndex:]
		}
		equation = strings.Trim(equation, "\n\t ")
		equations = append(equations, equation)
	}

	return equations, nil
}
