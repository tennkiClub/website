package database

import (
	"github.com/jinzhu/gorm"
	"model"
	//postgres connect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Profile OO Struct
type Profile struct {
	db *gorm.DB
}

//Update method of Profile
func (p *Profile) Update(profileNew *model.DataProfile) {
	var oldProfile model.DataProfile
	p.db.Where("Name = ?", profileNew.Name).First(&oldProfile)
	oldProfile.Icon = profileNew.Icon
	oldProfile.Job = profileNew.Job
	oldProfile.Content = profileNew.Content
	p.db.Save(&oldProfile)
}

//Add method of Profile
func (p *Profile) Add(profile *model.DataProfile) {
	p.db.Create(&profile)
}

//QueryAll method of Profile
func (p *Profile) QueryAll() *[]model.DataProfile {
	var profile []model.DataProfile
	p.db.Find(&profile)
	return &profile
}
