package utils

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
)

func DealErr(c *gin.Context, err error, log string) {
	c.JSON(500, gin.H{"code": 500, "msg": "服务器异常，请重试"})
	global.SugarLog.Error(log, err)
}
