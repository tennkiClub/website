package model

import (
	"github.com/jinzhu/gorm"
)

//Information struct
type DataInfo struct {
	gorm.Model
	URL     string `gorm:"unique"`
	Title   string
	Content string
}
