package public_api

import "github.com/gin-gonic/gin"

// Login 处理用户登录
func (PublicApi) Login(c *gin.Context) {
	//数据绑定

	//检查图形验证码

	//根据用户名密码查询数据库

	//生成token
	c.JSON(200, gin.H{"code": 200, "msg": "登录成功"})
}
