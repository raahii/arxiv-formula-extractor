package latex

import (
	"fmt"
	"log"
	"testing"
)

func TestRemoveComment(t *testing.T) {
	input := `%hogehoge
		% fugafuga
	\begin{equation}
	y=ax
	\end{equation}
	foo % bar
	`

	actual := RemoveComment(input)

	expected := `
		
	\begin{equation}
	y=ax
	\end{equation}
	foo 
	`
	if actual != expected {
		t.Fatalf("\ngot  %#v\nwant %#v", actual, expected)
	}
}

func TestFindMacros(t *testing.T) {
	input := `hogehoge
\def{\bmphi}{{\bm \phi}}
\newcommand{\bmphi}{{\bm \phi}}
\renewcommand{\set}[1]{\{#1\}}
\newcommand{\subf}[2]{%
  {\small\begin{tabular}[b]{@{}c@{}}
  #1\\#2
  \end{tabular}}%
}`

	actual, err := FindMacros(input)
	if err != nil {
		log.Fatal(err)
	}
	expected := []string{
		`\def{\bmphi}{{\bm \phi}}`,
		`\newcommand{\bmphi}{{\bm \phi}}`,
		`\renewcommand{\set}[1]{\{#1\}}`,
		`\newcommand{\subf}[2]{%` + "\n" +
			`  {\small\begin{tabular}[b]{@{}c@{}}` + "\n" +
			`  #1\\#2` + "\n" +
			`  \end{tabular}}%` + "\n" +
			`}`,
	}

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

func TestFindEquations(t *testing.T) {
	input := `hogehoge
	fugafuga
	\begin{equation}
	y=ax
	\end{equation}
	foobar
	\begin{align}
	y=ax
	\end{align}
	golang hogehoge
	\begin{eqnarray}
	y=ax
	\end{eqnarray}
	`

	actual := FindEquations(input)
	e := `y=ax`
	expected := []string{e, e, e}

	if len(actual) != len(expected) {
		t.Fatalf("number of elems is mismatch.\n\ngot %#v\nwant %#v", len(actual), len(expected))
	}

	for i := 0; i < len(actual); i++ {
		if actual[i] != expected[i] {
			t.Fatalf("%dth element\n\ngot  %#v\nwant %#v", i, actual[i], expected[i])
		}
	}
}
