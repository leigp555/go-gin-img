package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"img/server/global"
	"time"
)

func Log() gin.HandlerFunc {
	return func(c *gin.Context) {
		//生成请求id
		requestId := uuid.New().String()
		//传递给下面的中间件
		c.Set("requestId", requestId)
		//记录请求开始的时间
		startTime := time.Now()

		c.Next()

		//计算请求时间
		endTime := time.Now()
		latencyTime := endTime.Sub(startTime)

		//请求日志包含的要素
		reqMethod := c.Request.Method
		reqUri := c.Request.RequestURI
		reqClientIp := c.ClientIP()
		resStatusCode := c.Writer.Status()
		useAgent := c.Request.UserAgent()
		referer := c.Request.Referer()
		requestSize := c.Request.ContentLength
		responseSize := c.Writer.Size()
		global.Glog.Infof("\nrequest_id:%v,start_time:%v,latency:%v,method:%v,url:%v,status:%v,ip:%v,user_agent:%v,referer:%v,request_size:%v,response_size:%v\n", requestId, startTime, latencyTime, reqMethod, reqUri, resStatusCode, reqClientIp, useAgent, referer, requestSize, responseSize)
	}
}
