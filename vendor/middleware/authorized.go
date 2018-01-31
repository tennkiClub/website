package middleware

import (
	"github.com/gin-gonic/gin"
	"local_modules/githubauth"
	"net/http"
	"time"
)

//AuthContent struct
type AuthContent struct {
	OrgName  string
	ClientID string
}

//AuthGithubKey struct
type AuthGithubKey struct {
	ClientID  string
	SecretKey string
}

//AuthRequired func
func AuthRequired(authc *AuthContent) gin.HandlerFunc {
	return func(c *gin.Context) {
		githubLogin, _ := c.Request.Cookie("github_login")
		githubToken, _ := c.Request.Cookie("github_token")
		if githubLogin != nil && githubToken != nil {
			orgContent := githubauth.ContentProvider{Name: githubLogin.Value, Token: githubToken.Value, OrgName: authc.OrgName}
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

			keyContent := githubauth.ContentProvider{ClientID: authc.ClientID}
			c.Redirect(302, githubauth.GetGitHubAuth(&keyContent))
			c.Abort()
			return
		}

	}
}

//AuthGithubCallback middleware
func AuthGithubCallback(authk *AuthGithubKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		key := githubauth.ContentProvider{Code: code, ClientID: authk.ClientID, SecretKey: authk.SecretKey}
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
