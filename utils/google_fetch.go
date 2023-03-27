package utils

import (
	"encoding/json"
	"img/server/global"
	"net/http"
	"strconv"
	"time"
)

type f struct{}

var GoogleFetch = f{}

type GoogleUser struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (f) GoogleUerFetch(token string) (u GoogleUser, errMsg string, error error) {
	//使用access_token获取用户信息
	req, err := http.NewRequest("GET", "https://www.googleapis.com/drive/v2/files", nil)
	req.Header.Set("Authorization", "Bearer "+token)
	if err != nil {
		return GoogleUser{}, "请求创建失败", err
	}
	client := &http.Client{
		Timeout: time.Second * 12,
	}
	resp, err := client.Do(req)
	if err != nil {
		return GoogleUser{}, "获取用户信息失败", err
	}
	defer func() {
		if err := resp.Body.Close(); err != nil {
			global.SugarLog.Error("google user 响应体关闭失败", err)
		}
	}()
	var user = GoogleUser{}
	err = json.NewDecoder(resp.Body).Decode(&user)
	if err != nil {
		return GoogleUser{}, "github user响应体解析失败", err
	}
	return user, "", nil
}

// GoogleUserName 根据github返回的用户信息生成一个全新的用户名
func (f) GoogleUserName(u GoogleUser) string {
	return strconv.Itoa(u.ID) + "@google"
}

// GooglePassword 根据github返回的用户信息生成一个全新的密码
func (f) GooglePassword(u GoogleUser) string {
	return Md5Str(strconv.Itoa(u.ID))
}
