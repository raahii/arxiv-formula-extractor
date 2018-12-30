package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"strings"
)

type Feed struct {
	Link    Link    `xml:"href, attr"`
	Title   string  `xml:"title"`
	Id      string  `xml:"id"`
	Updated string  `xml:"updated"`
	Entries []Entry `xml:"entry"`
}

type Entry struct {
	Id        string   `xml:"id"`
	Updated   string   `xml:"updated"`
	Published string   `xml:"published"`
	Title     string   `xml:"title"`
	Summary   string   `xml:"summary"`
	Authors   []Author `xml:"author"`
	Category  Category `xml:"category"`
	Links     []Link   `xml:"link"`
}

type Category struct {
	Name string `xml:"term,attr"`
}

type Link struct {
	Title string `xml:"title,attr"`
	Value string `xml:"href,attr"`
	Type  string `xml:"type,attr"`
}

type Author struct {
	Name string `xml:"name"`
}

func parseXML(xmlStr string) Feed {
	var result Feed
	if err := xml.Unmarshal([]byte(xmlStr), &result); err != nil {
		log.Fatal(err)
	}
	return result
}

func searchPaper(titleQuery string) []Entry {
	u, err := url.Parse("http://export.arxiv.org/api/query")
	if err != nil {
		log.Fatal(err)
	}

	q := u.Query()
	q.Add("search_query", "ti:'"+titleQuery+"'")
	q.Add("sortBy", "lastUpdatedDate")
	q.Add("sortOrder", "ascending")
	q.Add("max_results", "5")
	u.RawQuery = q.Encode()

	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	xmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	xmlObj := parseXML(string(xmlBytes))

	return xmlObj.Entries
}

func EprintUrl(paper Entry) (string, error) {
	var url string
	var found bool

	found = false
	for i := 0; i < len(paper.Links); i++ {
		link := paper.Links[i]
		if link.Title == "pdf" {
			url = link.Value
			found = true
			break
		}
	}
	if !found {
		fmt.Errorf("The resource (e-print) of the requested paper is not found.")
	}

	url = strings.Replace(url, "/pdf", "/e-print", 1)
	url = strings.Replace(url, ".pdf", "", 1)
	return url, nil
}

func main() {
	entries := searchPaper("Generative Adversarial Nets")
	for i := 0; i < len(entries); i++ {
		fmt.Printf("- %+v\n", entries[i].Links)
	}
}
