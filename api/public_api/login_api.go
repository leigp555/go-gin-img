package public_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"strconv"
)

// Login 处理用户登录
func (PublicApi) Login(c *gin.Context) {
	//数据绑定
	type user struct {
		Username    string `form:"username" binding:"required,min=1,max=20" msg:"用户名不能为空,且长度为1~20位"`
		Password    string `form:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
		CaptchaId   string `form:"imgCaptchaId" binding:"required" msg:"图形验证码不正确"`
		CaptchaCode string `form:"imgCaptcha" binding:"required" msg:"图形验证码不正确"`
	}
	var userInfo user
	mdb := global.Mydb
	//验证json数据绑定
	err := c.ShouldBind(&userInfo)
	if err != nil {
		msg := utils.GetValidMsg(err, &userInfo)
		fmt.Println(err)
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{msg}}})
		return
	}
	//检查图形验证码
	isRight := utils.Captcha.Verify(userInfo.CaptchaId, userInfo.CaptchaCode)
	if !isRight {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"图形验证码错误"}}})
		return
	}
	//根据用户名密码查询数据库
	//验证用户名是否存在
	var u = models.User{}
	mdb.Where("username = ?", userInfo.Username).First(&u)
	if u.Username != userInfo.Username {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"用户名不存在"}}})
		return
	}
	//验证密码
	if u.Password != utils.Md5Str(userInfo.Password) {
		c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": []string{"密码错误"}}})
		return
	}
	//生成token
	token, err2 := utils.Token.Generate(strconv.Itoa(int(u.ID)))
	if err2 != nil {
		c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常,请稍后再试"}}})
		return
	}
	c.JSON(200, gin.H{"code": 200, "msg": "登录成功", "token": token, "user": map[string]any{"email": u.Email, "username": u.Username}})
}
