package utils

import (
	"fmt"
	"net/smtp"
)

type gmail struct{}

var GEmail = new(gmail)

func (gmail) Send(emailNumber string, randStr string) (err error) {
	// 设置邮件服务器信息
	smtpServer := "smtp.gmail.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", "xyleigp@gmail.com", "dyshkcrrlsdwyhhh", smtpServer)
	from := "IMG 图床"
	to := []string{emailNumber}
	subject := "IMG图床验证码"
	body := fmt.Sprintf(`
	  <p style="font-size:18px">您的验证码为: <span style="color:orange;font-weight:700">%s</span> 此验证码5分钟内有效</p>
    `, randStr)
	// 设置邮件内容
	msg := "From: " + from + "\r\n" +
		"To: " + to[0] + "\r\n" +
		"Subject: " + subject + "\r\n\r\n" +
		body
	// 发送邮件
	err = smtp.SendMail(smtpServer+":"+smtpPort, auth, from, []string{emailNumber}, []byte(msg))
	return err
}
