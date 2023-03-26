package public_api

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"time"
)

func (PublicApi) GithubLogin(c *gin.Context) {
	type GithubCallback struct {
		Code string `form:"code" binding:"required" msg:"code不能为空"`
	}
	var callback GithubCallback
	var Res = utils.Res
	var mdb = global.Mydb
	//绑定参数
	if err := c.ShouldBind(&callback); err != nil {
		msg := utils.GetValidMsg(err, &callback)
		Res.Fail.Normal(c, 400, msg)
		return
	}
	//获取GitHub Access Token
	token, msg, err := utils.GithubFetch.Token(callback.Code)
	if err != nil {
		Res.Fail.ErrorWithMsg(c, err, msg, msg)
		return
	}
	//获取GitHub用户信息
	u, errMsg, err := utils.GithubFetch.Uer(token)
	if err != nil || u.ID == 0 {
		Res.Fail.ErrorWithMsg(c, err, errMsg, errMsg)
		return
	}
	//数据库查询用户是否存在，没有就注册
	var user models.User
	var userEmail = utils.GithubFetch.Email(u)
	err = mdb.Where("email = ?", userEmail).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//注册用户
			username := utils.GithubFetch.UserName(u)
			password := utils.GithubFetch.Password(u)
			user = models.User{
				Email:    userEmail,
				Username: username,
				Password: password,
			}
			err = mdb.Create(&user).Error
			if err != nil {
				Res.Fail.ErrorWithMsg(c, err, "注册失败", "注册失败")
				return
			}
			goto Meta
		} else {
			Res.Fail.ErrorWithMsg(c, err, "获取用户信息失败", "获取用户信息失败")
		}
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
	responseContent := ResponseContent{
		Username:  user.Username,
		Email:     user.Email,
		Token:     token,
		CreatedAt: user.CreatedAt,
		UpdateAt:  user.UpdatedAt,
	}
	Res.Success.Normal(c, "登录成功", responseContent)
}
