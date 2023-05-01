package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
	"img/server/middleware"
)

type routerGroup struct{}

var rg = new(routerGroup)

// InitRouter 业务路由
func InitRouter(r *gin.Engine) {
	r.Use(middleware.Log())
	g := r.Group("v1/api")

	//1.静态资源托管
	g.Static("/static", "./public")

	//2.注册公共路由（访问不需要提供token）
	{
		rg.PublicRouter(g)
	}

	//3.私有路由（访问需要提供token）
	gr := g.Group("", middleware.TokenVerify())
	//用户信息相关的路由
	userGroup := gr.Group("/user")
	{
		rg.UserRouter(userGroup)
	}
	//图片相关的路由
	imageGroup := gr.Group("/img")
	{
		rg.ImgRouter(imageGroup)
	}
}

// InitSocketRouter socket 路由
func InitSocketRouter(r *gin.Engine) {
	r.GET("/chat", api.GroupApi.SocketApi.Chat)
}
