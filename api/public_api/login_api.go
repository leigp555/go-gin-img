package public_api

import "github.com/gin-gonic/gin"

func (PublicApi) Login(c *gin.Context) {
	c.JSON(200, gin.H{"code": 200, "msg": "登录成功"})
}
