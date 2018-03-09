package middleware

import (
	"github.com/gin-gonic/gin"
	"local_modules/configloader"
	"local_modules/githubauth"
	"net/http"
	"time"
)

//AuthRequired func
func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := configloader.New("config.yaml")
		githubLogin, _ := c.Request.Cookie("github_login")
		githubToken, _ := c.Request.Cookie("github_token")
		if githubLogin != nil && githubToken != nil {
			orgContent := githubauth.ContentProvider{Name: githubLogin.Value, Token: githubToken.Value, OrgName: config.Github.OrgName}
			code := githubauth.GetOrg(&orgContent)
			if code != 200 {
				expiration := time.Unix(0, 0)
				logincookie := http.Cookie{Name: "github_login", Value: "", Path: "/", Expires: expiration}
				tokencookie := http.Cookie{Name: "github_token", Value: "", Path: "/", Expires: expiration}
				http.SetCookie(c.Writer, &logincookie)
				http.SetCookie(c.Writer, &tokencookie)
				c.HTML(http.StatusOK, "admin/winfo", gin.H{
					"title": "驗證失敗",
					"info":  "使用者不存在",
				})
				c.Abort()
				return
			}
			c.Next()
		} else {

			keyContent := githubauth.ContentProvider{ClientID: config.Github.ClientID}
			c.Redirect(302, githubauth.GetGitHubAuth(&keyContent))
			c.Abort()
			return
		}

	}
}

//AuthGithubCallback middleware
func AuthGithubCallback() gin.HandlerFunc {
	return func(c *gin.Context) {
		config := configloader.New("config.yaml")
		code := c.Query("code")
		key := githubauth.ContentProvider{Code: code, ClientID: config.Github.ClientID, SecretKey: config.Github.SecretKey}
		accesstoken := githubauth.GetToken(&key)
		username := githubauth.GetUsername(accesstoken)
		expiration := time.Now().Add(1 * time.Hour)
		logincookie := http.Cookie{Name: "github_login", Value: username, Path: "/", Expires: expiration}
		tokencookie := http.Cookie{Name: "github_token", Value: accesstoken, Path: "/", Expires: expiration}
		http.SetCookie(c.Writer, &logincookie)
		http.SetCookie(c.Writer, &tokencookie)
		c.Next()
	}
}
