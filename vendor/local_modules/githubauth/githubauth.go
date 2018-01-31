package githubauth

import (
	"bytes"
	"github.com/parnurzeal/gorequest"
)

//github base url
var githubBase = "https://github.com/login/oauth/authorize"

//ContentProvider struct
type ContentProvider struct {
	Code      string
	ClientID  string
	SecretKey string
	OrgName   string
	Token     string
	Name      string
}

//GetGitHubAuth func
func GetGitHubAuth(c *ContentProvider) string {
	var buffer bytes.Buffer
	buffer.WriteString(githubBase)
	buffer.WriteString("?client_id=")
	buffer.WriteString(c.ClientID)
	buffer.WriteString("&scope=user%20admin:org%20repo&allow_singup=false")
	return buffer.String()
}

//GetToken export  func
func GetToken(c *ContentProvider) string {
	type jsonAccessToken struct {
		AccessToken string `json:"access_token"`
		TokenType   string
		Scope       string
	}
	type Data struct {
		ClientID     string `json:"client_id"`
		ClientSecret string `json:"client_secret"`
		Code         string `json:"code"`
	}
	var jsonToken jsonAccessToken
	dataSend := Data{ClientID: c.ClientID, ClientSecret: c.SecretKey, Code: c.Code}
	gorequest.New().Post("https://github.com/login/oauth/access_token/").Set("Accept", "application/json").Send(dataSend).EndStruct(&jsonToken)
	return jsonToken.AccessToken
}

//GetUsername func
func GetUsername(token string) string {
	type jsonUsername struct {
		Login string
	}
	var tokenbuffer bytes.Buffer
	var jsonUser jsonUsername
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(token)
	gorequest.New().Get("https://api.github.com/user").Set("Authorization", tokenbuffer.String()).EndStruct(&jsonUser)
	return jsonUser.Login
}

//GetOrg func
func GetOrg(c *ContentProvider) int {
	var tokenbuffer bytes.Buffer
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(c.Token)
	var urlbuffer bytes.Buffer
	urlbuffer.WriteString("https://api.github.com/orgs/")
	urlbuffer.WriteString(c.OrgName)
	urlbuffer.WriteString("/memberships/")
	urlbuffer.WriteString(c.Name)
	resp, _, _ := gorequest.New().Get(urlbuffer.String()).Set("Authorization", tokenbuffer.String()).End()
	return resp.StatusCode
}
