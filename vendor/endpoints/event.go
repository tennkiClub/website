package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Event View
func Event(c *gin.Context) {
	c.HTML(http.StatusOK, "main/event", gin.H{
		"title": "index",
	})
}
