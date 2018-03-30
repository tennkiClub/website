package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"local_modules/configloader"
	"model"
	//postgres connect
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

//Service strut
type Service struct {
	db          *gorm.DB
	Information *Information
	Profile     *Profile
}

//New function of Create Service
func New() *Service {
	c := configloader.New("config.yaml")
	connString := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=%s password=%s", c.DB.Host, c.DB.User, c.DB.DBName, c.DB.SSLMode, c.DB.Password)
	fmt.Println(connString)
	db, err := gorm.Open("postgres", connString)
	if err != nil {
		panic("failed to connect database")
	}
	info := &Information{
		db: db,
	}
	profile := &Profile{
		db: db,
	}
	return &Service{
		db:          db,
		Information: info,
		Profile:     profile,
	}
}

//CreateSchema func of Service
func (s *Service) CreateSchema() {
	s.db.AutoMigrate(&model.DataInfo{})
	s.db.AutoMigrate(&model.DataProfile{})
}

//CloseDBConnect func of service
func (s *Service) CloseDBConnect() {
	defer s.db.Close()
}
