package arxiv

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
	Text string
}
