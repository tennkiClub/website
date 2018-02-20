package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//Profile Endpoint
func Profile(c *gin.Context) {
	c.HTML(http.StatusOK, "main/profile", gin.H{
		"title": "Profile",
	})
}
