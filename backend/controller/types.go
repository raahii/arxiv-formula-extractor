package controller

import (
	// "github.com/jinzhu/gorm"
	"time"
)

type Model struct {
	ID        uint       `json:"id" gorm:"primary_key"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
	DeletedAt *time.Time `json:"deleted_at"`
}

type Paper struct {
	Model
	ArxivId    string     `json:"arxiv_id gorm:"type:varchar(20);not null;unique_index"`
	Authors    []Author   `json:"authors" gorm:"not null;association_save_reference:true;foreignkey:PaperID"`
	Title      string     `json:"title" gorm:"not null"`
	Abstract   string     `json:"abstract" gorm:"not null;type:text"`
	AbsUrl     string     `json:"url" gorm:"not null"`
	TarballUrl string     `json:"tarball_url" gorm:"not null"`
	Macros     string     `json:"macros" gorm:"type:varchar(10000)"`
	Equations  []Equation `json:"equations" gorm:"association_save_reference:true;foreignkey:PaperID"`
}

func (Paper) TableName() string {
	return "papers"
}

type Author struct {
	Model
	Name    string `json:"name" gorm:not null`
	PaperID uint   `json:"paper_id" gorm:"not null"`
}

func (Author) TableName() string {
	return "authors"
}

type Equation struct {
	Model
	Expression string `json:"expression" gorm:"not null;type:varchar(10000)"`
	Body       string `json:"body" gorm:"not null;type:varchar(10000)"`
	PaperID    uint   `json:"paper_id" gorm:"not null"`
}

func (Equation) TableName() string {
	return "equations"
}
