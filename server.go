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
			dbtest.CreateSchema()
			rand.Seed(time.Now().UnixNano())
			testinfo := model.DataInfo{URL: strconv.Itoa(rand.Intn(65535)), Title: "test", Content: "testit"}
			dbtest.Information.Add(&testinfo)
			dbtest.Information.QueryLimit(10)
			testprofile := model.DataProfile{Icon: "/static/img/icon/mhh.jpg", Name: "棉花", Content: "所有的形象圖,社團角色,插圖以及東方電腦合同之主要執筆者<br/>熟悉插畫/漫畫以及各種不同風格之作畫 <br />上至百合下至BL 都難不倒他", Job: "繪師"}
			//testprofile := model.DataProfile{Icon: "/static/img/icon/0mu.jpg", Name: "零夢", Content: "沒什麼反應就是一位攤主(? <br /> 偶爾幹蠢事被社團的各位嗆 <br /> 平時負責網站維護,印刷,跟別攤裝逼（？", Job: "攤主/被嗆的"}
			dbtest.Profile.Update(&testprofile)
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
