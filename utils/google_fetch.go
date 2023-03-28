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
		ID          string `json:"permissionId"`
		Email       string `json:"emailAddress"`
		Kind        string `json:"kind"`
		DisplayName string `json:"displayName"`
		PhotoLink   string `json:"photoUrl"`
		Me          bool   `json:"me"`
	} `json:"user"`
}

func (f) GoogleUerFetch(token string) (u GoogleUser, errMsg string, error error) {
	//使用access_token获取用户信息
	url := fmt.Sprintf("https://www.googleapis.com/drive/v3/about?access_token=%s&fields=user", token)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return GoogleUser{}, "请求创建失败", err
	}
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return GoogleUser{}, "Google拒绝了请求", err
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		// 处理错误
		return
	}
	fmt.Println("==========================")
	fmt.Println(string(body))
	fmt.Println("==========================")
	defer func() {
		if err := resp.Body.Close(); err != nil {
			global.SugarLog.Error("google user 响应体关闭失败", err)
		}
	}()
	var user = GoogleUser{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		fmt.Println("===================")
		fmt.Println(err)
		fmt.Println(user)
		return GoogleUser{}, "github user响应体解析失败", err
	}
	return user, "", nil
}

// GoogleUserName 根据github返回的用户信息生成一个全新的用户名
func (f) GoogleUserName(u GoogleUser) string {
	return u.User.ID + "@google"
}

// GooglePassword 根据github返回的用户信息生成一个全新的密码
func (f) GooglePassword(u GoogleUser) string {
	return Md5Str(u.User.ID)
}
