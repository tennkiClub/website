package githubauth

import (
	"bytes"
	"encoding/json"
	"net/http"
)

//github base url
var githubBase = "https://github.com/login/oauth/authorize"

//Key Service
type Key struct {
	Code      string
	ClientID  string
	SecretKey string
}

//OrgContent strct
type OrgContent struct {
	Name    string
	Token   string
	OrgName string
}

//JSONAccessToken struct
type jsonAccessToken struct {
	AccessToken string `json:"access_token"`
	TokenType   string
	Scope       string
}

type jsonUsername struct {
	Login string
}

//GetGitHubAuth func
func GetGitHubAuth(key *Key) string {
	var buffer bytes.Buffer
	buffer.WriteString(githubBase)
	buffer.WriteString("?client_id=")
	buffer.WriteString(key.ClientID)
	buffer.WriteString("&scope=user%20admin:org%20repo&allow_singup=false")
	return buffer.String()
}

//GetToken export  func
func GetToken(key *Key) string {
	data := "{\"client_id\":\"" + key.ClientID +
		"\", \"client_secret\":\"" + key.SecretKey +
		"\", \"code\":\"" + key.Code +
		"\"}"
	var jsonStr = []byte(data)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var jsonToken jsonAccessToken
	err = decoder.Decode(&jsonToken)
	if err != nil {
		panic(err)
	}
	return jsonToken.AccessToken
}

//GetUsername func
func GetUsername(token string) string {
	var tokenbuffer bytes.Buffer
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(token)
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	req.Header.Set("Authorization", tokenbuffer.String())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	decoder := json.NewDecoder(resp.Body)
	var jsonUser jsonUsername
	err = decoder.Decode(&jsonUser)
	if err != nil {
		panic(err)
	}
	return jsonUser.Login
}

//GetOrg func
func GetOrg(orgc *OrgContent) int {
	var tokenbuffer bytes.Buffer
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(orgc.Token)
	var urlbuffer bytes.Buffer
	urlbuffer.WriteString("https://api.github.com/orgs/")
	urlbuffer.WriteString(orgc.OrgName)
	urlbuffer.WriteString("/memberships/")
	urlbuffer.WriteString(orgc.Name)
	req, err := http.NewRequest("GET", urlbuffer.String(), nil)
	req.Header.Set("Authorization", tokenbuffer.String())
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	status := resp.StatusCode
	return status
}
