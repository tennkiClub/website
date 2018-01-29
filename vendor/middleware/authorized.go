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
			orgContent := githubauth.OrgContent{githubLogin.Value, githubToken.Value, authc.OrgName}
			code := githubauth.GetOrg(&orgContent)
			if code != 200 {
				c.Redirect(302, "/error/usernotfound")
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

//AuthGithubCallback middleware
func AuthGithubCallback(authk *AuthGithubKey) gin.HandlerFunc {
	return func(c *gin.Context) {
		code := c.Query("code")
		key := githubauth.Key{code, authk.ClientID, authk.SecretKey}
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
