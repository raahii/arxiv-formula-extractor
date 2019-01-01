package handlers

import (
	"time"
)

type Paper struct {
	Id        int32      `json:"id"`
	Version   int32      `json:"version:`
	Authors   string     `json:"authors"`
	Title     string     `json:"title"`
	Abstract  string     `json:"abstract"`
	Equations []Equation `json:"equations"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

// type Entry struct {
// 	Id        string   `xml:"id" json:"id"`
// 	Updated   string   `xml:"updated" json:"updated"`
// 	Published string   `xml:"published" json:"published"`
// 	Title     string   `xml:"title" json:"title"`
// 	Summary   string   `xml:"summary" json:"summary"`
// 	Authors   []Author `xml:"author" json:"authors"`
// 	Category  Category `xml:"category" json:"category"`
// 	Links     []Link   `xml:"link" json:"links"`
// }

type Equation struct {
	Name string `json:name`
}
