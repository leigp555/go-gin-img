package utils

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
)

//统一Response

type response struct {
	Success Success
	Fail    Fail
}
type Success struct{}
type Fail struct{}

var Res = response{}

//成功相关的响应

func (s Success) Normal(c *gin.Context, msg string, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(200, gin.H{"success": true, "code": 200, "msg": msg, "requestId": requestId, "data": data})
}
func (s Success) WidthData(c *gin.Context, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(200, gin.H{"success": true, "code": 200, "msg": "success", "requestId": requestId, "data": data})
}
func (s Success) WidthMsg(c *gin.Context, msg string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(200, gin.H{"success": true, "code": 200, "msg": msg, "requestId": requestId})
}

//失败相关的响应

func (s Fail) Normal(c *gin.Context, code int, msg string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(code, gin.H{"success": false, "code": code, "msg": msg, "requestId": requestId})
}
func (s Fail) WidthData(c *gin.Context, code int, msg string, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(code, gin.H{"success": false, "code": code, "msg": msg, "requestId": requestId, "data": data})
}
func (s Fail) Error(c *gin.Context, err error, log string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	global.Slog.Error(log, err)
	c.JSON(500, gin.H{"success": false, "code": 500, "msg": "服务繁忙,请重试", "requestId": requestId})
}

func (s Fail) ErrorWithMsg(c *gin.Context, err error, log string, msg string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	global.Slog.Error(log, err)
	c.JSON(500, gin.H{"success": false, "code": 500, "msg": msg, "requestId": requestId})
}
