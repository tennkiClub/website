// Package main provides ...
package main

import (
	"database"
	"endpoints"
	"fmt"
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"math/rand"
	"middleware"
	"model"
	"os"
	"strconv"
	"time"
)

func main() {
	if len(os.Args) > 1 {
		if os.Args[1] == "--runserver" {
			// gin.SetMode(gin.ReleaseMode)
			route := gin.Default()
			route.Static("/static", "./static")
			render := eztemplate.New()
			render.TemplatesDir = "template/"
			render.Ext = ".tmpl"
			render.Debug = true
			route.HTMLRender = render.Init()
			// route for Index
			route.GET("/", endpoints.Index)
			route.GET("/test", endpoints.Test)
			route.GET("/profile", endpoints.Profile)
			route.GET("/oauth/callback", middleware.AuthGithubCallback(), endpoints.OauthCallback)
			route.GET("/admin", middleware.AuthRequired(), endpoints.AdminIndex)
			route.GET("/info/:url", endpoints.Information)
			route.Run()
		} else if os.Args[1] == "--dbinit" {
			dbinit := database.New()
			dbinit.CreateSchema()
			dbinit.CloseDBConnect()
		} else if os.Args[1] == "--dbtest" {
			dbtest := database.New()
			rand.Seed(time.Now().UnixNano())
			testinfo := model.DataInfo{URL: strconv.Itoa(rand.Intn(65535)), Title: "test", Content: "testit"}
			dbtest.Information.Add(&testinfo)
			dbtest.Information.QueryLimit(10)
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
