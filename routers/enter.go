package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
	"img/server/middleware"
)

type routerGroup struct{}

var router = new(routerGroup)

// InitRouter 业务路由
func InitRouter(r *gin.Engine) {
	gp := r.Group("v1/api", middleware.Log())

	//1.静态资源托管
	gp.Static("/static", "./public")

	//2.注册公共路由（访问不需要提供token）
	{
		router.public(gp)
	}

	//3.私有路由（访问需要提供token）
	gs := gp.Group("", middleware.TokenVerify())
	//用户信息相关的路由
	{
		router.user(gs.Group("/user"))
	}
	//图片相关的路由
	{
		router.img(gs.Group("/img"))
	}
}

// InitSocketRouter socket 路由
func InitSocketRouter(r *gin.Engine) {
	r.GET("/chat", api.Group.Socket.Chat)
}
