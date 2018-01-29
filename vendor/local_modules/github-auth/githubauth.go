package githubauth

import (
	"bytes"
	"io/ioutil"
	"net/http"
)

//github base url
var githubBase = "https://github.com/login/oauth/authorize"

//GetGitHubAuth func
func GetGitHubAuth(clientID string) string {
	var buffer bytes.Buffer
	buffer.WriteString(githubBase)
	buffer.WriteString("?client_id")
	buffer.WriteString(clientID)
	buffer.WriteString("scope=user%20admin:org%20repo&allow_singup=false")
	return buffer.String()
}

//GetToken export  func
func GetToken(code string, clientID string, secretKey string) []byte {
	data := "{\"client_id\":\"" + clientID +
		"\", \"client_secret\":\"" + secretKey +
		"\", \"code\":\"" + code +
		"\"}"
	var jsonStr = []byte(data)
	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token/", bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

//GetUsername func
func GetUsername(token string) []byte {
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
	body, _ := ioutil.ReadAll(resp.Body)
	return body
}

//GetOrg func
func GetOrg(name string, token string) int {
	var tokenbuffer bytes.Buffer
	tokenbuffer.WriteString("token ")
	tokenbuffer.WriteString(token)
	var urlbuffer bytes.Buffer
	urlbuffer.WriteString("https://api.github.com/orgs/")
	urlbuffer.WriteString(orgname)
	urlbuffer.WriteString("/memberships/")
	urlbuffer.WriteString(name)
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
