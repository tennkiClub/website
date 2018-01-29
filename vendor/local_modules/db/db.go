package db

import (
	"github.com/jinzhu/gorm"
	//postgres connect
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

//CreateSchema func of Service
func (s *Service) CreateSchema() {
	s.db.AutoMigrate(&Information{})
}

//UpdateInformation func of Service
func (s *Service) UpdateInformation(infoNew *Information) {
	var oldInfo Information
	s.db.Where("URL = ?", infoNew.URL).First(&oldInfo)
	oldInfo.Title = infoNew.Title
	oldInfo.Content = infoNew.Content
	s.db.Save(&oldInfo)
}

//AddInformation func of Service
func (s *Service) AddInformation(info *Information) {
	s.db.Create(&info)
}

//QueryLimitInformation func of Service
func (s *Service) QueryLimitInformation(count int) *[]Information {
	var info []Information
	s.db.Limit(10).Find(&info)
	return &info
}

//QueryOneInformation func of Service
func (s *Service) QueryOneInformation(url string) *Information {
	var info Information
	s.db.Where("URL = ?", url).First(&info)
	return &info
}

//CloseDBConnect func of service
func (s *Service) CloseDBConnect() {
	defer s.db.Close()
}
