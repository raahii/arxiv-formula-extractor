package handlers

import (
	"time"
)

type Paper struct {
	Id        int32      `json:"id"`
	Version   int32      `json:"version"`
	Authors   string     `json:"authors"`
	Title     string     `json:"title"`
	Abstract  string     `json:"abstract"`
	Equations []Equation `json:"equations"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

type Equation struct {
	Name string `json:name`
}
