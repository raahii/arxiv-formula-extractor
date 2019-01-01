package arxiv

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
	Id        string   `xml:"id" json:"id"`
	Updated   string   `xml:"updated" json:"updated"`
	Published string   `xml:"published" json:"published"`
	Title     string   `xml:"title" json:"title"`
	Summary   string   `xml:"summary" json:"summary"`
	Authors   []Author `xml:"author" json:"author"`
	Category  Category `xml:"category" json:"category"`
	Links     []Link   `xml:"link" json:"link"`
}

type Category struct {
	Name string `xml:"term,attr" json:"name"`
}

type Link struct {
	Title string `xml:"title,attr"`
	Value string `xml:"href,attr"`
	Type  string `xml:"type,attr"`
}

type Author struct {
	Name string `xml:"name"`
}

type Equation struct {
	Text string
}
