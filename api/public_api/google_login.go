package public_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"strconv"
	"time"
)

func (PublicApi) GoogleLogin(c *gin.Context) {
	type GoogleLogin struct {
		AccessToken string `form:"access_token" binding:"required" msg:"access_token不能为空"`
	}
	var googleLogin GoogleLogin
	var mdb = global.Mdb
	//绑定参数
	if err := c.ShouldBind(&googleLogin); err != nil {
		msg := utils.GetValidMsg(err, &googleLogin)
		utils.Res.Fail(c, 400, msg, struct{}{})
		return
	}
	//获取用户信息
	u, msg, err := utils.GoogleFetch.GoogleUerFetch(googleLogin.AccessToken)
	if err != nil || u.User.PermissionId == "" {
		utils.Res.FailWidthRecord(c, 500, msg, struct{}{}, err, "获取用户信息失败")
		return
	}
	//数据库查询用户是否存在，没有就注册
	var user models.User
	var username = utils.GoogleFetch.GoogleUserName(u)
	err = mdb.Where("username = ?", username).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//注册用户
			password := utils.GoogleFetch.GooglePassword(u)
			//获取uid
			uid, err := utils.GetUid()
			if err != nil {
				utils.Res.FailWidthRecord(c, 500, "注册失败,请重试", struct{}{}, err, "uid生成失败")
				return
			}
			user = models.User{
				Username:       username,
				Password:       password,
				NickName:       utils.GenerateNackName(),
				Uid:            uid,
				DiskLimit:      5120,
				DiskUsage:      0,
				Role:           "normal",
				LastLogin:      time.Now().Format("2006-01-02 15:04:05"),
				LastLoginIp:    c.ClientIP(),
				CurrentLogin:   time.Now().Format("2006-01-02 15:04:05"),
				CurrentLoginIp: c.ClientIP(),
			}
			err = mdb.Create(&user).Error
			if err != nil {
				utils.Res.FailWidthRecord(c, 500, "注册失败", struct{}{}, err, "用户注册数据写入mysql失败")
				return
			}
			goto Meta
		} else {
			utils.Res.FailWidthRecord(c, 500, "注册失败", struct{}{}, err, "用户注册数据查询mysql失败")
		}
		return
	}
	//已经注册,直接登录
	//修改登录时间和IP地址
	user.LastLogin = user.CurrentLogin
	user.LastLoginIp = user.CurrentLoginIp
	user.CurrentLogin = time.Now().Format("2006-01-02 15:04:05")
	user.CurrentLoginIp = c.ClientIP()
	if err := mdb.Save(&user).Error; err != nil {
		utils.Res.FailWidthRecord(c, 500, "登录失败,请重试", struct{}{}, err, "mysql更新失败")
		return
	}
Meta:
	type ResponseContent struct {
		Username  string    `json:"username"`
		Email     string    `json:"email"`
		Token     string    `json:"token"`
		CreatedAt time.Time `json:"createdAt"`
		UpdateAt  time.Time `json:"updateAt"`
	}

	//生成token
	token, err := utils.Token.Generate(strconv.Itoa(int(user.ID)))
	if err != nil {
		utils.Res.FailWidthRecord(c, 500, "注册失败", struct{}{}, err, "token生成失败")
		return
	}
	responseContent := ResponseContent{
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdatedAt,
	}
	utils.Res.Success(c, "登录成功", responseContent)
}
