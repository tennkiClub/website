// Package main provides ...
package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	eztemplate "github.com/michelloworld/ez-gin-template"
	"os"
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
			route.GET("/", Index)
			route.Run()
		} else if os.Args[1] == "--dbinit" {
			dbinit := New()
			dbinit.createSchema()
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
