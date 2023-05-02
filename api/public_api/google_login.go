package public_api

import (
	"fmt"
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
	var userEmail = u.User.EmailAddress
	err = mdb.Where("email = ?", userEmail).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			//注册用户
			username := utils.GoogleFetch.GoogleUserName(u)
			password := utils.GoogleFetch.GooglePassword(u)
			user = models.User{
				Email:    userEmail,
				Username: username,
				Password: password,
			}
			err = mdb.Create(&user).Error
			if err != nil {
				fmt.Println("===========================")
				fmt.Println(err)
				utils.Res.FailWidthRecord(c, 500, "注册失败", struct{}{}, err, "用户注册数据写入mysql失败")
				return
			}
			goto Meta
		} else {
			utils.Res.FailWidthRecord(c, 500, "注册失败", struct{}{}, err, "用户注册数据查询mysql失败")
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
