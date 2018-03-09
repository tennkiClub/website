package endpoints

import (
	"database"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Information View
func Information(c *gin.Context) {
	url := c.Param("url")
	db := database.New()
	info := db.Information.QueryOne(url)
	c.HTML(http.StatusOK, "main/info", gin.H{
		"title": info.Title,
		"info":  info,
	})
}
