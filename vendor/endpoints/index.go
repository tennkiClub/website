package endpoints

import (
	"database"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Index View
func Index(c *gin.Context) {
	db := database.New()
	info := db.Information.QueryLimit(10)
	c.HTML(http.StatusOK, "main/index", gin.H{
		"title": "index",
		"info":  info,
	})
}

//Test View
func Test(c *gin.Context) {
	c.String(http.StatusOK, "Hello")
}
