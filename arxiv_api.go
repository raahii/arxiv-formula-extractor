package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

type Feed struct {
	Link    string  `xml:"link"`
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
	Author    []Author `xml:"author"`
	Category  string   `xml:"category"`
}

type Author struct {
	Name string `xml:"name"`
}

type Paper struct {
	Id   string `xml:"id"`
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

func main() {
	entries := searchPaper("Generative Adversarial Nets")
	for i := 0; i < len(entries); i++ {
		fmt.Printf("- %+v\n", entries[i])
	}
}
