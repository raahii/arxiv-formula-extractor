package latex

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
