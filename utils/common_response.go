package utils

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
)

//统一Response

type Res struct {
	Success Success
	Fail    Fail
}
type Success struct{}
type Fail struct{}

func (s Success) Normal(c *gin.Context, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	c.JSON(200, gin.H{"code": 200, "msg": "success", "requestId": requestId, "data": data})
}

func (s Success) WidthMsg(c *gin.Context, msg string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	c.JSON(200, gin.H{"code": 200, "msg": msg, "requestId": requestId})
}

func (s Fail) Normal(c *gin.Context, code int, msg string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	c.JSON(code, gin.H{"code": code, "msg": msg, "requestId": requestId})
}
func (s Fail) WidthData(c *gin.Context, code int, msg string, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	c.JSON(code, gin.H{"code": code, "msg": msg, "requestId": requestId, "data": data})
}
func (s Fail) Error(c *gin.Context) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.SugarLog.Warn("上下文获取requestId失败")
	}
	c.JSON(500, gin.H{"code": 500, "msg": "服务繁忙,请重试", "requestId": requestId})
}
