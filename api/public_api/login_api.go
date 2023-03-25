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
	//验证json数据绑定
	if err := c.ShouldBind(&userInfo); err != nil {
		msg := utils.GetValidMsg(err, &userInfo)
		c.JSON(400, gin.H{"code": 400, "msg": msg})
		return
	}
	//检查图形验证码
	isRight := utils.Captcha.Verify(userInfo.CaptchaId, userInfo.CaptchaCode)
	if !isRight {
		c.JSON(400, gin.H{"code": 400, "msg": "验证码不正确"})
		return
	}
	//根据用户名密码查询数据库
	var u = models.User{}
	if err := mdb.Where("username = ?", userInfo.Username).First(&u).Error; err != nil {
		utils.DealErr(c, err, "/public_api/login 登录用户查询失败")
		return
	}
	if u.Username != userInfo.Username {
		c.JSON(400, gin.H{"code": 400, "msg": "用户名不存在,请先注册"})
		return
	}
	//验证密码
	if u.Password != utils.Md5Str(userInfo.Password) {
		c.JSON(400, gin.H{"code": 400, "msg": "密码错误"})
		return
	}
	//生成token
	token, err := utils.Token.Generate(strconv.Itoa(int(u.ID)))
	if err != nil {
		utils.DealErr(c, err, "/public_api/login 生成token失败")
		return
	}
	type ResponseContent struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
		RequestId string    `json:"requestId"`
		UpdateAt  time.Time `json:"updateAt"`
	}
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	responseContent := ResponseContent{
		Username:  u.Username,
		Email:     u.Email,
		Token:     token,
		RequestId: requestId.(string),
		UpdateAt:  u.UpdatedAt,
	}
	c.JSON(200, gin.H{"code": 200, "msg": "登录成功", "data": responseContent})
}
