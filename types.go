package main

type Feed struct {
	Link         Link    `xml:"href, attr"`
	Title        string  `xml:"title"`
	Id           string  `xml:"id"`
	Updated      string  `xml:"updated"`
	Entries      []Entry `xml:"entry"`
	TotalResults int32   `xml:"totalResults"`
	StartIndex   int32   `xml:"startIndex"`
	ItemsPerPage int32   `xml:"itemsPerPage"`
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
