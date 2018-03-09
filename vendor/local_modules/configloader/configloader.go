package configloader

import (
	"fmt"
	"github.com/jinzhu/configor"
)

// Config Struct
type Config struct {
	APPName string
	Github  struct {
		ClientID  string
		SecretKey string
		OrgName   string
	}
	DB struct {
		Host     string
		User     string
		DBName   string
		SSLMode  string
		Password string
	}
}

//New Func
func New(c string) *Config {
	var Conf Config
	configor.Load(&Conf, c)
	fmt.Printf("config: %#v", Conf.APPName)
	return &Conf
}
