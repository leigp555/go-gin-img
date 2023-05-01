package utils

import (
	"encoding/json"
	"img/server/global"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

type fetch struct{}

var GithubFetch = fetch{}

func (fetch) Token(code string) (token string, errMsg string, error error) {
	//根据前端code 获取access_token
	data := url.Values{}
	data.Set("client_id", global.Config.Login.Github.ClientId)
	data.Set("client_secret", global.Config.Login.Github.ClientSecret)
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
			global.Slog.Error("github access_token响应体关闭失败")
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

func (fetch) Uer(token string) (u User, errMsg string, error error) {
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
			global.Slog.Error("github user 响应体关闭失败", err)
		}
	}()
	var user = User{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return User{}, "github user响应体解析失败", err
	}
	return user, "", nil
}

// Email 根据github返回的用户信息生成一个全新的邮箱
func (fetch) Email(u User) string {
	return strconv.Itoa(u.ID) + "@github.com"
}

func (fetch) UserName(u User) string {
	return u.Login + "@github"
}
func (fetch) Password(u User) string {
	return Md5Str(strconv.Itoa(u.ID))
}
