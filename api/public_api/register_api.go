package public_api

import (
	"context"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"time"
)

// Register 处理用户注册
func (PublicApi) Register(c *gin.Context) {
	type NewUserInfo struct {
		Username         string `form:"username" binding:"required,alphanum,min=1,max=20" msg:"用户名不能为空,且长度为1~20位数字或字母"`
		Email            string `form:"email" binding:"required,email" msg:"请输入正确的邮箱"`
		EmailCaptchaCode string `form:"emailCaptcha" binding:"required,len=6" msg:"邮箱验证码不正确"`
		Password         string `form:"password" binding:"required,min=6,max=12,alphanum" msg:"密码不能为空,且长度为6~12位数字或字母"`
		RePassword       string `form:"rePassword" binding:"required,min=6,max=12,eqfield=Password" msg:"两次输入的密码不一致"`
		ImgCaptchaId     string `form:"imgCaptchaId" binding:"required" msg:"图形验证码ID不能为空"`
		ImgCaptcha       string `form:"imgCaptcha" binding:"required,numeric,len=4" msg:"图形验证码不正确"`
	}

	var newUserInfo NewUserInfo
	rdb := global.Redb
	mdb := global.Mydb
	//验证数据绑定
	err := c.ShouldBind(&newUserInfo)
	if err != nil {
		msg := utils.GetValidMsg(err, &newUserInfo)
		c.JSON(400, gin.H{"code": 400, "msg": msg})
		return
	}

	//验证图形验证码
	b := utils.Captcha.Verify(newUserInfo.ImgCaptchaId, newUserInfo.ImgCaptcha)
	if b == false {
		c.JSON(400, gin.H{"code": 400, "msg": "图形验证码错误"})
		return
	}
	//验证email验证码
	var ctx = context.Background()
	val, err := rdb.Get(ctx, newUserInfo.Email).Result()
	if err != nil {
		utils.DealErr(c, err, "redis获取数据失败")
		return
	}
	if val != newUserInfo.EmailCaptchaCode {
		c.JSON(400, gin.H{"code": 400, "msg": "邮箱验证码错误"})
		return
	}
	//查询邮箱或者用户名是否已经被使用
	var u = models.User{}
	if err = mdb.Where("email = ?", newUserInfo.Email).Or("username=?", newUserInfo.Username).Find(&u).Error; err != nil {
		utils.DealErr(c, err, "邮箱查询数据库失败")
		return
	}
	if u.Email == newUserInfo.Email {
		c.JSON(400, gin.H{"code": 400, "msg": "邮箱已绑定"})
		return
	}
	//查询用户名是否已存在
	if u.Username == newUserInfo.Username {
		c.JSON(400, gin.H{"code": 400, "msg": "用户名已存在"})
		return
	}
	//将密码md5
	s := utils.Md5Str(newUserInfo.Password)
	newUserInfo.Password = s
	//保存到数据库
	newUser := models.User{
		Username: newUserInfo.Username,
		Email:    newUserInfo.Email,
		Password: newUserInfo.Password,
	}
	if err = mdb.Create(&newUser).Error; err != nil {
		utils.DealErr(c, err, "写入数据库失败，用户注册不成功")
		return
	}
	//注册成功，返回客户端
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	type ResponseContent struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		RequestId string    `json:"requestId"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
	responseContent := ResponseContent{
		Username:  newUser.Username,
		Email:     newUser.Email,
		RequestId: requestId.(string),
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	c.JSON(200, gin.H{"code": 200, "msg": "注册成功", "data": responseContent})
}
