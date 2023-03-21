package public_api

import (
	"github.com/gin-gonic/gin"
	"img/server/utils"
)

// Register 处理用户注册
func (PublicApi) Register(c *gin.Context) {
	type NewUserInfo struct {
		Username         string `form:"username" binding:"required,min=1,max=20" msg:"用户名不能为空,且长度为1~20位"`
		Email            string `form:"email" binding:"required,email" msg:"请输入正确的邮箱"`
		EmailCaptchaCode string `form:"emailCaptcha" binding:"required,len=6" msg:"邮箱验证码不正确"`
		Password         string `form:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
		RePassword       string `form:"rePassword" binding:"required,min=6,max=12,eqfield=Password" msg:"两次输入的密码不一致"`
		ImgCaptchaId     string `form:"imgCaptchaId" binding:"required" msg:"图形验证码不正确"`
		ImgCaptcha       string `form:"imgCaptcha" binding:"required" msg:"图形验证码不正确"`
	}

	var newUserInfo NewUserInfo
	err := c.ShouldBind(&newUserInfo)
	if err != nil {
		msg := utils.GetValidMsg(err, &newUserInfo)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": msg}})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "注册成功", "body": newUserInfo})

}
