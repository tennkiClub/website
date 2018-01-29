package middleware

import (
	"github.com/gin-gonic/gin"
	"local_modules/githubauth"
	"net/http"
)

//AuthContent struct
type AuthContent struct {
	OrgName  string
	ClientID string
}

//AuthRequired func
func AuthRequired(authc *AuthContent) gin.HandlerFunc {
	return func(c *gin.Context) {
		githubLogin, _ := c.Request.Cookie("github_login")
		githubToken, _ := c.Request.Cookie("github_token")
		if githubLogin != nil && githubToken != nil {
			orgContent := githubauth.OrgContent{githubLogin.Value, githubToken.Value, authc.OrgName}
			code := githubauth.GetOrg(&orgContent)
			if code != 200 {
				c.HTML(http.StatusOK, "main/info", gin.H{
					"title": "驗證失敗",
					"info":  "使用者不存在",
				})
				return
			}
			c.Next()
		} else {

			keyContent := githubauth.Key{ClientID: authc.ClientID}
			c.Redirect(302, githubauth.GetGitHubAuth(&keyContent))
			return
		}

	}
}
