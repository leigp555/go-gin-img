package utils

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
)

//统一Response

type response struct{}

var Res = response{}

type responseBody struct {
	Success   bool   `json:"success"`
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      any    `json:"data"`
	RequestId string `json:"requestId"`
}

//成功相关的响应

func (response) Success(c *gin.Context, msg string, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(200, responseBody{Success: true, Code: 200, Msg: msg, RequestId: requestId.(string), Data: data})
}

//失败相关的响应

func (response) Fail(c *gin.Context, code int, msg string, data any) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(code, responseBody{Success: false, Code: code, Msg: msg, RequestId: requestId.(string), Data: data})
}

func (response) FailWidthRecord(c *gin.Context, code int, msg string, data any, err any, errMsg string) {
	requestId, exist := c.Get("requestId")
	if !exist {
		requestId = "0-0-0-0"
		global.Slog.Warn("上下文获取requestId失败")
	}
	c.JSON(code, responseBody{Success: false, Code: code, Msg: msg, RequestId: requestId.(string), Data: data})
	global.Slog.Errorf("%s err: %s", errMsg, err)
}
