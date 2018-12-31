package main

import (
	"testing"
)

func TestRemoveTags(t *testing.T) {
	input := `\begin{equation}
	\label{eq:minimaxgame-definition}
	y=ax
	\end{equation}`

	actual := RemoveTagLines(input, []string{`\label`})

	expected := `\begin{equation}
	y=ax
	\end{equation}`
	if actual != expected {
		t.Fatalf("\ngot  %#v\nwant %#v", actual, expected)
	}
}

func TestFindEquations(t *testing.T) {
	input := `hogehoge
	fugafuga
	\begin{equation}
	y=ax
	\end{equation}
	foobar
	`

	actual := FindEquations(input)[0]

	expected := Equation{`y=ax`}
	if actual != expected {
		t.Fatalf("\ngot  %#v\nwant %#v", actual, expected)
	}
}
