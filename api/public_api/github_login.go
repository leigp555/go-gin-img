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

func (PublicApi) GithubLogin(c *gin.Context) {
	type GithubCallback struct {
		Code string `form:"code" binding:"required" msg:"code不能为空"`
	}
	var callback GithubCallback
	var mdb = global.Mdb
	//绑定参数
	if err := c.ShouldBind(&callback); err != nil {
		msg := utils.GetValidMsg(err, &callback)
		utils.Res.Fail(c, 400, msg, struct{}{})
		return
	}
	//获取GitHub Access Token
	githubToken, msg, err := utils.GithubFetch.Token(callback.Code)
	if err != nil {
		utils.Res.FailWidthRecord(c, 500, msg, struct{}{}, err, "获取GitHub Access Token失败")
		return
	}
	//获取GitHub用户信息
	u, errMsg, err := utils.GithubFetch.Uer(githubToken)
	if err != nil || u.ID == 0 {
		utils.Res.FailWidthRecord(c, 500, errMsg, struct{}{}, err, "获取GitHub用户信息失败")
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
