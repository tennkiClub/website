// Package main provides ...
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"local_modules/db"
	"math/rand"
	"os"
	"strconv"
	"time"
	"views"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--runserver" {
			route := gin.Default()
			route.Static("/static", "./static")
			render := eztemplate.New()
			render.TemplatesDir = "template/"
			render.Ext = ".tmpl"
			render.Debug = true
			route.HTMLRender = render.Init()
			// route for Index
			route.GET("/", view.Index)
			route.GET("/info/:url", view.Information)
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
