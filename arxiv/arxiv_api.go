package arxiv

import (
	"encoding/xml"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
	"os"
	"strings"
)

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

func SearchPapers(titleQuery string) Feed {
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

	return xmlObj
}
