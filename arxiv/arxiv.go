package arxiv

import (
	"encoding/xml"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
)

type Feed struct {
	Link         Link    `xml:"href, attr" json:"link"`
	Title        string  `xml:"title" json:"title"`
	Id           string  `xml:"id" json:"id"`
	Updated      string  `xml:"updated" json:"updated"`
	Entries      []Entry `xml:"entry" json:"entries"`
	TotalResults int32   `xml:"totalResults" json:"total_results"`
	StartIndex   int32   `xml:"startIndex" json:"start_index"`
	ItemsPerPage int32   `xml:"itemsPerPage" json:"items_per_page"`
}

type Entry struct {
	Id        string   `xml:"id" json:"id"`
	Updated   string   `xml:"updated" json:"updated"`
	Published string   `xml:"published" json:"published"`
	Title     string   `xml:"title" json:"title"`
	Summary   string   `xml:"summary" json:"summary"`
	Authors   []Author `xml:"author" json:"authors"`
	Category  Category `xml:"category" json:"category"`
	Links     []Link   `xml:"link" json:"links"`
}

type Category struct {
	Name string `xml:"term,attr" json:"name"`
}

type Link struct {
	Title string `xml:"title,attr" json:"title"`
	Value string `xml:"href,attr" json:"value"`
	Type  string `xml:"type,attr" json:"type"`
}

type Author struct {
	Name string `xml:"name" json:"name"`
}

type Equation struct {
	Text string `json:"name"`
}

func DownloadTarball(url string, path string) error {
	// Create the file
	out, err := os.Create(path)
	if err != nil {
		return err
	}
	defer out.Close()

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}

	return nil
}

func parseXML(xmlStr string) Feed {
	var result Feed
	if err := xml.Unmarshal([]byte(xmlStr), &result); err != nil {
		log.Fatal(err)
	}

	return result
}

func SearchPapers(params map[string]string) (Feed, error) {
	// define api url
	u, err := url.Parse("http://export.arxiv.org/api/query")
	if err != nil {
		return Feed{}, err
	}

	// construct query string
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	u.RawQuery = q.Encode()

	// send the request
	resp, err := http.Get(u.String())
	if err != nil {
		return Feed{}, err
	}

	// parse result xml
	xmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Feed{}, err
	}
	xmlObj := parseXML(string(xmlBytes))

	return xmlObj, nil
}
