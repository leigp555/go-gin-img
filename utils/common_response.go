package utils

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
)

//统一Response

func Res(c *gin.Context, code int, data any, msg any) {
	c.JSON(code, gin.H{
		"code": code,
		"data": data,
		"msg":  msg})
}

func ResWidthMsg(c *gin.Context, code int, data any, msg any) {
	getMsg, ok := ErrMap[code]
	if ok {
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  getMsg})
	} else {
		global.SugarLog.Warnf("%v状态码是未知的状态码", code)
		c.JSON(code, gin.H{
			"code": code,
			"data": data,
			"msg":  "未知错误"})
	}

}
