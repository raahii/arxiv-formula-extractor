package main

import (
	"fmt"
	"os"
	"testing"
)

func Exists(filename string) bool {
	_, err := os.Stat(filename)
	return err == nil
}

func Remove(path string) {
	if Exists(path) {
		if err := os.Remove(path); err != nil {
			fmt.Println(err)
		}
	}
}

func TestDownloadTarball(t *testing.T) {
	sampleUrl := "https://dev-files.blender.org/file/download/bwdp5reejwpkuh5i2oak/PHID-FILE-nui3bpuan4wdvd7yzjrs/sample.tar.gz"
	saveTo := "/tmp/sample.tar.gz"
	err := DownloadTarball(sampleUrl, saveTo)
	Remove(saveTo)
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
}

func TestSearchPaper(t *testing.T) {
	result := SearchPaper("Generative Adversarial Nets")

	if result.TotalResults <= 0 {
		t.Fatalf("total numbers of results must be more than 0")
	}
}

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
