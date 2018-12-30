package main

import (
	"fmt"
	"os"
	"testing"
)

func TestEprintUrl(t *testing.T) {
	paper := Entry{}
	link := Link{"pdf", "http://arxiv.org/pdf/astro-ph/0608371v1.pdf", ""}
	paper.Links = []Link{link}

	actual, err := EprintUrl(paper)
	if err != nil {
		t.Fatalf("The function returned error %#v", err)
	}

	expected := "http://arxiv.org/e-print/astro-ph/0608371v1"
	if actual != expected {
		t.Fatalf("\ngot %#v\nwant %#v", actual, expected)
	}
}
