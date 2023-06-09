package test

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
	"testing"
)

func send(emailNumber []string, randStr string) (err error) {
	e := email.NewEmail()
	//设置发送方的邮箱
	e.From = "907090585@qq.com"
	// 设置接收方的邮箱
	e.To = emailNumber
	//设置主题
	e.Subject = "Email"
	//设置文件发送的内容

	htmlStr := fmt.Sprintf(`
    <h1>验证码</h1>
    <p>验证码为%s<p>
    <p>验证码仅在5分钟内有效<p>
    `, randStr)
	e.HTML = []byte(htmlStr)
	//设置服务器相关的配置
	err = e.Send("smtp.qq.com:25", smtp.PlainAuth("", "907090585@qq.com", "cvxjaeubymkxbbic", "smtp.qq.com"))
	return
}

func TestSendEmail(t *testing.T) {
	err := send([]string{"122974945@qq.com"}, "123456")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("success")
}
