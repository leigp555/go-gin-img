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
	rdb := global.Rdb
	mdb := global.Mdb
	//验证数据绑定
	if err := c.ShouldBind(&newUserInfo); err != nil {
		msg := utils.GetValidMsg(err, &newUserInfo)
		utils.Res.Fail(c, 400, msg, struct{}{})
		return
	}

	//验证图形验证码
	b := utils.Captcha.Verify(newUserInfo.ImgCaptchaId, newUserInfo.ImgCaptcha)
	if b == false {
		utils.Res.Fail(c, 400, "图形验证码错误", struct{}{})
		return
	}
	//验证email验证码
	var ctx = context.Background()
	val, err := rdb.Get(ctx, newUserInfo.Email).Result()
	if err != nil || val != newUserInfo.EmailCaptchaCode {
		utils.Res.Fail(c, 400, "邮箱验证码错误", struct{}{})
		return
	}
	//查询邮箱或者用户名是否已经被使用
	var u = models.User{}
	mdb.Where("email = ?", newUserInfo.Email).Or("username=?", newUserInfo.Username).Find(&u)
	if u.Email == newUserInfo.Email {
		utils.Res.Fail(c, 400, "邮箱已绑定", struct{}{})
		return
	}
	//查询用户名是否已存在
	if u.Username == newUserInfo.Username {
		utils.Res.Fail(c, 400, "用户名已存在", struct{}{})
		return
	}
	//获取uid
	uid, err := utils.GetUid()
	if err != nil {
		utils.Res.FailWidthRecord(c, 500, "注册失败,请重试", struct{}{}, err, "uid生成失败")
		return
	}
	//将密码md5
	s := utils.Md5Str(newUserInfo.Password)
	newUserInfo.Password = s
	//保存到数据库
	newUser := models.User{
		Username:       newUserInfo.Username,
		Email:          newUserInfo.Email,
		Password:       newUserInfo.Password,
		NickName:       utils.GenerateNackName(),
		Uid:            uid,
		DiskLimit:      5120,
		DiskUsage:      0,
		Role:           "normal",
		EmailVerified:  true,
		LastLogin:      time.Now().Format("2006-01-02 15:04:05"),
		LastLoginIp:    c.ClientIP(),
		CurrentLogin:   time.Now().Format("2006-01-02 15:04:05"),
		CurrentLoginIp: c.ClientIP(),
	}
	if err = mdb.Create(&newUser).Error; err != nil {
		utils.Res.FailWidthRecord(c, 500, "注册失败,请重试", struct{}{}, err, "用户注册数据存储失败")
		return
	}
	//注册成功，返回客户端
	type ResponseContent struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		CreatedAt time.Time `json:"createdAt"`
		UpdatedAt time.Time `json:"updatedAt"`
	}
	responseContent := ResponseContent{
		Username:  newUser.Username,
		Email:     newUser.Email,
		CreatedAt: newUser.CreatedAt,
		UpdatedAt: newUser.UpdatedAt,
	}
	utils.Res.Success(c, "注册成功", responseContent)
}
