package endpoints

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//AuthError view
func AuthError(c *gin.Context) {
	c.HTML(http.StatusOK, "admin/winfo", gin.H{
		"title": "權限錯誤",
		"info":  "使用者不存在",
	})

}
