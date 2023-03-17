package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/middleware"
)

func InitRouter(r *gin.Engine) {
	//添加全局跨域中间件
	r.Use(middleware.Cors())
	//配置路由路口
	g := r.Group("v1/api")
	{
		g.GET("/xxx", func(c *gin.Context) {
			c.JSON(200, gin.H{"hell": "world"})
		})
	}
}
