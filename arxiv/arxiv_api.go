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

func SearchPapers(params map[string]string) Feed {
	// define api url
	u, err := url.Parse("http://export.arxiv.org/api/query")
	if err != nil {
		log.Fatal(err)
	}

	// construct query string
	q := u.Query()
	for k, v := range params {
		q.Add(k, v)
	}
	// q.Add("sortBy", "lastUpdatedDate")
	// q.Add("sortOrder", "ascending")
	// q.Add("max_results", "5")
	u.RawQuery = q.Encode()

	// send the request
	resp, err := http.Get(u.String())
	if err != nil {
		log.Fatal(err)
	}

	// parse result xml
	xmlBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	xmlObj := parseXML(string(xmlBytes))

	return xmlObj
}
