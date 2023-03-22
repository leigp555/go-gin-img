package public_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
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
		ImgCaptchaId     string `form:"imgCaptchaId" binding:"required" msg:"图形验证码ID未获取"`
		ImgCaptcha       string `form:"imgCaptcha" binding:"required" msg:"图形验证码未获取"`
	}

	var newUserInfo NewUserInfo
	rdb := global.Redb
	mdb := global.Mydb

	//验证数据绑定
	err := c.ShouldBind(&newUserInfo)
	if err != nil {
		msg := utils.GetValidMsg(err, &newUserInfo)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": msg}})
		return
	}
	//验证图形验证码
	b := utils.Captcha.Verify(newUserInfo.ImgCaptchaId, newUserInfo.ImgCaptcha)
	if b == false {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"验证码错误"}}})
		return
	}
	//验证email验证码
	var ctx = context.Background()
	val, err := rdb.Get(ctx, newUserInfo.Email).Result()
	if err != nil || val != newUserInfo.EmailCaptchaCode {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": "邮箱验证码不正确"}})
		return
	}
	//查询邮箱是否已经被注册
	var u = models.User{}
	mdb.Where("email = ?", newUserInfo.Email).Find(&u)
	if u.Email == newUserInfo.Email {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": "邮箱已被绑定"}})
		return
	}
	//查询用户名是否已存在
	if u.Username == newUserInfo.Username {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": "用户已存在"}})
		return
	}
	//将密码md5
	s := utils.Md5Str(newUserInfo.Password)
	newUserInfo.Password = s
	//注册成功，返回客户端
	c.JSON(200, gin.H{"code": 200, "msg": "注册成功", "body": newUserInfo})
}
