package public_api

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"strconv"
	"time"
)

// Login 处理用户登录
func (PublicApi) Login(c *gin.Context) {
	//数据绑定
	type user struct {
		Username    string `form:"username" binding:"required,alphanum,min=1,max=20" msg:"用户名不能为空,且长度为1~20位数字或字母"`
		Password    string `form:"password" binding:"required,min=6,max=12,alphanum" msg:"密码不能为空,且长度为6~12位数字或字母"`
		CaptchaId   string `form:"imgCaptchaId" binding:"required" msg:"验证码id不正确"`
		CaptchaCode string `form:"imgCaptcha" binding:"required,numeric,len=4" msg:"验证码不正确"`
	}
	var userInfo user
	mdb := global.Mydb
	res := utils.Res
	//验证json数据绑定
	if err := c.ShouldBind(&userInfo); err != nil {
		msg := utils.GetValidMsg(err, &userInfo)
		res.Fail.Normal(c, 400, msg)
		return
	}
	//检查图形验证码
	isRight := utils.Captcha.Verify(userInfo.CaptchaId, userInfo.CaptchaCode)
	if !isRight {
		res.Fail.Normal(c, 400, "验证码不正确")
		return
	}
	//根据用户名密码查询数据库
	var u = models.User{}
	if err := mdb.Where("username = ?", userInfo.Username).First(&u).Error; err != nil {
		res.Fail.Normal(c, 400, "用户名不存在,请先注册")
		return
	}
	//验证密码
	if u.Password != utils.Md5Str(userInfo.Password) {
		res.Fail.Normal(c, 400, "密码错误")
		return
	}
	//生成token
	token, err := utils.Token.Generate(strconv.Itoa(int(u.ID)))
	if err != nil {
		res.Fail.Error(c, err, "/public_api/login 生成token失败")
		return
	}
	type ResponseContent struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"createdAt"`
		UpdateAt  time.Time `json:"updateAt"`
	}
	responseContent := ResponseContent{
		Username:  u.Username,
		Email:     u.Email,
		Token:     token,
		CreatedAt: u.CreatedAt,
		UpdateAt:  u.UpdatedAt,
	}
	res.Success.Normal(c, "登录成功", responseContent)
}
