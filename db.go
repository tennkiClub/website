package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Service strut
type Service struct {
	db *gorm.DB
}

//Information strut
type Information struct {
	gorm.Model
	URL     string `gorm:"unique"`
	Title   string
	Content string
}

//New function of Create Service
func New() *Service {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=website sslmode=disable password=password")
	if err != nil {
		panic("failed to connect database")
	}
	return &Service{
		db: db,
	}
}

func (s *Service) createSchema() {
	s.db.Debug().AutoMigrate(&Information{})
	s.db.Close()
}

func (s *Service) updateInformation(infoNew *Information) {
	var oldInfo Information
	s.db.Where("URL = ?", infoNew.URL).First(&oldInfo)
	oldInfo.Title = infoNew.Title
	oldInfo.Content = infoNew.Content
	s.db.Save(&oldInfo)
	s.db.Close()
}

func (s *Service) addInformation(info *Information) {
	s.db.Create(&info)
	s.db.Close()
}
