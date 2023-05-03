package utils

import (
	"fmt"
	"net/smtp"
)

type qemail struct{}

var QEmail = new(qemail)

func (qemail) Send(emailNumber string, randStr string) (err error) {
	// 设置邮件服务器信息
	smtpServer := "smtp.qq.com"
	smtpPort := "587"
	auth := smtp.PlainAuth("", "2026499232@qq.com", "ndzqxcvuekkkdjhd", smtpServer)
	from := "2026499232@qq.com"
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
	fmt.Println([]string{emailNumber})
	return err
}
