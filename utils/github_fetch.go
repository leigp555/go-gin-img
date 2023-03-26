package utils

import (
	"encoding/json"
	"img/server/global"
	"net/http"
	"net/url"
	"strings"
)

type Fetch struct{}

var GithubFetch = Fetch{}

func (Fetch) Token(code string) (token string, errMsg string, error error) {
	//根据前端code 获取access_token
	data := url.Values{}
	data.Set("client_id", "97bc323362f96abbc3d3")
	data.Set("client_secret", "f4c82475993cebf0a259750f537aec3d17004dcc")
	data.Set("code", code)

	req, err := http.NewRequest("POST", "https://github.com/login/oauth/access_token", strings.NewReader(data.Encode()))
	if err != nil {
		return "", "构造github access_token请求失败", err
	}
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	req.Header.Set("Accept", "application/json")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", "获取github access_token失败", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			global.SugarLog.Error("github access_token响应体关闭失败")
		}
	}()
	var result struct {
		AccessToken string `json:"access_token"`
		TokenType   string `json:"token_type"`
		Scope       string `json:"scope"`
	}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return "", "github access_token响应体解析失败", err
	}
	return result.AccessToken, "", nil
}

type User struct {
	ID    int    `json:"id"`
	Login string `json:"login"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func (Fetch) Uer(token string) (u User, errMsg string, error error) {
	//使用access_token获取用户信息
	req, err := http.NewRequest("GET", "https://api.github.com/user", nil)
	if err != nil {
		return User{}, "构造github user请求失败", err
	}
	req.Header.Set("Authorization", "token "+token)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return User{}, "获取github user失败", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			global.SugarLog.Error("github user 响应体关闭失败", err)
		}
	}()
	var user = User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return User{}, "github user响应体解析失败", err
	}
	return user, "", nil
}
