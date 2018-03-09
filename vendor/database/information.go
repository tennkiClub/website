package database

import (
	"github.com/jinzhu/gorm"
	"model"
	//postgres connect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Information struct
type Information struct {
	db *gorm.DB
}

//Update method of Information
func (i *Information) Update(infoNew *model.DataInfo) {
	var oldInfo model.DataInfo
	i.db.Where("URL = ?", infoNew.URL).First(&oldInfo)
	oldInfo.Title = infoNew.Title
	oldInfo.Content = infoNew.Content
	i.db.Save(&oldInfo)
}

func (i *Information) Add(info *model.DataInfo) {
	i.db.Create(&info)
}

func (i *Information) QueryLimit(count int) *[]model.DataInfo {
	var info []model.DataInfo
	i.db.Limit(count).Find(&info)
	return &info
}

func (i *Information) QueryOne(url string) *model.DataInfo {
	var info model.DataInfo
	i.db.Where("URL = ?", url).First(&info)
	return &info
}
