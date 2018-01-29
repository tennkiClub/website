package view

import (
	"github.com/gin-gonic/gin"
	"local_modules/db"
	"net/http"
)

//Index View
func Index(c *gin.Context) {
	database := db.New()
	info := database.QueryLimitInformation(10)
	c.HTML(http.StatusOK, "main/index", gin.H{
		"title": "index",
		"info":  info,
	})
}
