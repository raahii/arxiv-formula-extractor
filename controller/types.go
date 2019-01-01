package controller

import (
	"time"
)

type Paper struct {
	Id        int32      `json:"id" gorm:"primary_key;not null;unique"`
	Version   int32      `json:"version" gorm:"not null"`
	Authors   string     `json:"authors" gorm:"not null"`
	Title     string     `json:"title" gorm:"not null"`
	Abstract  string     `json:"abstract" gorm:"not null"`
	Equations []Equation `json:"equations"`
	UpdatedAt time.Time  `json:"updated_at"`
	CreatedAt time.Time  `json:"created_at"`
}

func (Paper) TableName() string {
	return "papers"
}

type Equation struct {
	Id   int    `json:id gorm:"primary_key;not null;AUTO_INCREMENT"`
	Name string `json:name gorm:"not null"`
}

func (Equation) TableName() string {
	return "equations"
}
