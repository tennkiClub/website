package endpoints

import (
	"database"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

//Profile Endpoint
func Profile(c *gin.Context) {
	db := database.New()
	profile := db.Profile.QueryAll()
	fmt.Println(profile)
	c.HTML(http.StatusOK, "main/profile", gin.H{
		"title":   "Profile",
		"profile": profile,
	})
}
