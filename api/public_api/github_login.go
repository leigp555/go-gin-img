package public_api

import (
	"github.com/gin-gonic/gin"
	"img/server/utils"
)

func (PublicApi) GithubLogin(c *gin.Context) {
	type GithubCallback struct {
		Code string `form:"code" binding:"required" msg:"code不能为空"`
	}
	var callback GithubCallback
	var Res = utils.Res
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
	if err != nil {
		Res.Fail.ErrorWithMsg(c, err, errMsg, errMsg)
		return
	}

	Res.Success.WidthData(c, u)
}
