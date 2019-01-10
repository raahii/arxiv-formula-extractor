package latex

import (
	"fmt"
	"log"
	"testing"
)

func assertArray(actual []string, expected []string, t *testing.T) {
	if len(actual) != len(expected) {
		msg := "number of elements is mismatch!\n\n"
		msg += fmt.Sprintf("got %d elements\nwant %d elements\n\n", len(actual), len(expected))
		msg += fmt.Sprintf("got %#v\nwant %#v\n\n", actual, expected)
		t.Fatalf(msg)
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%dth element is mismatch!\n\ngot  %#v\nwant %#v", i, actual[i], expected[i])
		}
	}
}

func TestRemoveComment(t *testing.T) {
	input := `%hogehoge
		% fugafuga
	\begin{equation}
	y=ax
	\end{equation}
	foo % bar
	aiueo %
	`

	actual := RemoveComment(input)

	expected := `
		
	\begin{equation}
	y=ax
	\end{equation}
	foo 
	aiueo 	`

	if actual != expected {
		t.Fatalf("\ngot  %#v\nwant %#v", actual, expected)
	}
}

// func TestRemoveSimpleCommands(t *testing.T) {
// 	input := `\label{hogehoge}
// 	\hogehoge\label{aiu}
// 	\label{aa\hoge{foovar}}
// 	\foo{}
// 	\foobar{hoge}`
//
// 	actual, err := RemoveSimpleCommands(input, []string{`\label`, `\foo`})
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	expected := `
// 	\hogehoge
//
//
// 	\foobar{hoge}`
// 	if actual != expected {
// 		t.Fatalf("\ngot  %#v\nwant %#v", actual, expected)
// 	}
// }

func TestFindParenthesis(t *testing.T) {
	input := `{{\small\begin{tabular}[b]{@{}c@{}}#1\\#2\end{tabular}}}`

	endIndex, err := FindParenthesis(input, '{')
	if err != nil {
		log.Fatal(err)
	}
	actual := input[0:endIndex]

	expected := `{{\small\begin{tabular}[b]{@{}c@{}}#1\\#2\end{tabular}}}`

	if actual != expected {
		t.Fatalf("\ngot  %#v\nwant %#v", actual, expected)
	}
}

func TestFindMacroCommands(t *testing.T) {
	input := `hogehoge
\def{\bmphi}{{\bm \phi}}
\newcommand{\bmphi}{{\bm \phi}}
\renewcommand{\set}[1]{\{#1\}}
\newcommand{\subf}[2]{%
  {\small\begin{tabular}[b]{@{}c@{}}
  #1\\#2
  \end{tabular}}%
}`

	actual, err := FindMacroCommands(RemoveComment(input))
	if err != nil {
		log.Fatal(err)
	}

	expected := []string{
		`\def{\bmphi}{{\bm \phi}}`,
		`\newcommand{\bmphi}{{\bm \phi}}`,
		`\renewcommand{\set}[1]{\{#1\}}`,
		`\newcommand{\subf}[2]{` +
			`  {\small\begin{tabular}[b]{@{}c@{}}` + "\n" +
			`  #1\\#2` + "\n" +
			`  \end{tabular}}}`,
	}

	assertArray(actual, expected, t)
}

func TestFindEquations(t *testing.T) {
	input := `hogehoge
	fugafuga
	\begin{equation}
	y=ax
	\end{equation}
	\begin{equation*}
	y=ax
	\end{equation*}
	foobar
	\begin{aligned}
	y=ax
	\end{aligned}
	golang hogehoge
	\begin{eqnarray}
	y=ax
	\end{eqnarray}
	\begin{eqnarray}
	y=ax
	\end{eqnarray}
	`

	actual, err := FindEquations(input)
	if err != nil {
		log.Fatal(err)
	}

	e := `y=ax`
	expected := []string{e, e, e, e}

	assertArray(actual, expected, t)
}
