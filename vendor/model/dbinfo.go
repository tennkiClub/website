package model

import (
	"github.com/jinzhu/gorm"
	"html/template"
)

//DataInfo model struct
type DataInfo struct {
	gorm.Model
	URL     string `gorm:"unique"`
	Title   string
	Content string
}

//DataProfile model struct
type DataProfile struct {
	gorm.Model
	Icon    string
	Name    string `gorm:"unique"`
	Job     string
	Content template.HTML
}
