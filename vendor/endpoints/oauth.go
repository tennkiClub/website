package endpoints

import (
	"github.com/gin-gonic/gin"
)

//OauthCallback endpoints
func OauthCallback(c *gin.Context) {
	c.Redirect(302, "/admin")
}
