package middleware

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"strconv"
	"strings"
)

func TokenVerify() gin.HandlerFunc {
	return func(c *gin.Context) {
		var mdb = global.Mdb
		//从请求头获取token
		tokenHeader := c.GetHeader("Authorization")
		//从请求头获取token失败
		if tokenHeader == "" {
			utils.Res.Fail(c, 401, "请上传身份凭证", struct{}{})
			c.Abort()
			return
		}
		//拆分出token
		splitArr := strings.Split(tokenHeader, " ")
		tokenStr := splitArr[1]
		//解析token 解析失败阻止后续中间件执行
		userId, err2 := utils.Token.Parse(tokenStr)
		if err2 != nil {
			utils.Res.Fail(c, 403, "用户身份过期,请重新登录", struct{}{})
			c.Abort()
			return
		}
		var dbUser = models.User{}
		mdb.Where("id = ?", userId).First(&dbUser)
		//将用户的信息传给其余的中间件
		c.Set("userId", strconv.Itoa(int(dbUser.ID)))
		c.Set("userEmail", dbUser.Email)
	}
}
