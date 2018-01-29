package endpoints

import (
	"github.com/gin-gonic/gin"
	"local_modules/db"
	"net/http"
)

//Information View
func Information(c *gin.Context) {
	url := c.Param("url")
	database := db.New()
	info := database.QueryOneInformation(url)
	c.HTML(http.StatusOK, "main/info", gin.H{
		"title": info.Title,
		"info":  info,
	})
}
