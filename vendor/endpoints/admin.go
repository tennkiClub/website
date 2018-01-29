package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//AdminIndex view
func AdminIndex(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/index", gin.H{
		"title": "管理界面",
	})

}
