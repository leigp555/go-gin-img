package utils

import (
	"encoding/json"
	"fmt"
	"img/server/global"
	"io/ioutil"
	"net/http"
)

type f struct{}

var GoogleFetch = f{}

type GoogleUser struct {
	User struct {
		Kind         string `json:"kind"`
		DisplayName  string `json:"displayName"`
		PhotoLink    string `json:"photoLink"`
		Me           bool   `json:"me"`
		PermissionId string `json:"permissionId"`
		EmailAddress string `json:"emailAddress"`
	} `json:"user"`
}

func (f) GoogleUerFetch(token string) (u GoogleUser, errMsg string, error error) {
	//使用access_token获取用户信息
	url := fmt.Sprintf("https://www.googleapis.com/drive/v3/about?access_token=%s&fields=user", token)
	fmt.Println(url)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GoogleUser{}, "请求创建失败", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GoogleUser{}, "Google拒绝了请求", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			global.SugarLog.Error("google user 响应体关闭失败", err)
		}
	}()
	jsonBytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		return GoogleUser{}, "读取Google用户信息失败", err
	}
	var response GoogleUser
	err = json.Unmarshal(jsonBytes, &response)
	if err != nil {
		// 处理错误
		return GoogleUser{}, "解析Google用户信息失败", err
	}
	return response, "", nil
}

// GoogleUserName 根据github返回的用户信息生成一个全新的用户名
func (f) GoogleUserName(u GoogleUser) string {
	return u.User.DisplayName + "@google"
}

// GooglePassword 根据github返回的用户信息生成一个全新的密码
func (f) GooglePassword(u GoogleUser) string {
	return Md5Str(u.User.PermissionId)
}
