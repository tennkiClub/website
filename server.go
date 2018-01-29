// Package main provides ...
package main

import (
	"endpoints"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/configor"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"local_modules/db"
	"math/rand"
	"middleware"
	"os"
	"strconv"
	"time"
)

//Config struct
var Config = struct {
	APPName string
	Github  struct {
		ClientID  string
		SecretKey string
		OrgName   string
	}
}{}

func main() {
	configor.Load(&Config, "config.yaml")
	if len(os.Args) > 1 {
		if os.Args[1] == "--runserver" {
			authc := middleware.AuthContent{Config.Github.OrgName, Config.Github.ClientID}
			route := gin.Default()
			route.Static("/static", "./static")
			render := eztemplate.New()
			render.TemplatesDir = "template/"
			render.Ext = ".tmpl"
			render.Debug = true
			route.HTMLRender = render.Init()
			// route for Index
			route.GET("/", endpoints.Index)
			route.GET("/admin", middleware.AuthRequired(&authc), endpoints.AdminIndex)
			route.GET("/info/:url", endpoints.Information)
			route.Run()
		} else if os.Args[1] == "--dbinit" {
			dbinit := db.New()
			dbinit.CreateSchema()
			dbinit.CloseDBConnect()
		} else if os.Args[1] == "--dbtest" {
			dbtest := db.New()
			rand.Seed(time.Now().UnixNano())
			testinfo := db.Information{URL: strconv.Itoa(rand.Intn(65535)), Title: "test", Content: "testit"}
			dbtest.AddInformation(&testinfo)
			dbtest.QueryLimitInformation(10)
			dbtest.CloseDBConnect()
		} else {
			fmt.Println("Please use:")
			fmt.Println("--runserver : run gin Server. ")
			fmt.Println("--dbinit : run postgreSQL init.  ")
			os.Exit(3)
		}
	} else {
		fmt.Println("Please use:")
		fmt.Println("--runserver : run gin Server. ")
		fmt.Println("--dbinit : run postgreSQL init.  ")
		os.Exit(3)
	}
}
