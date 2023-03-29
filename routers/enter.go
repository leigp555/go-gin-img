package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/middleware"
)

type ApiRouterGroup struct{}

var apiRouterGroup = new(ApiRouterGroup)

func InitRouter(r *gin.Engine) {
	//添加全局跨域中间件
	r.Use(middleware.Cors(), middleware.Log())
	//配置路由路口
	g := r.Group("v1/api")

	//静态资源托管
	r.Static("/static", "./public")

	//注册公共路由（访问不需要提供token）
	{
		apiRouterGroup.PublicRouter(g)
	}

	//私有路由（访问需要提供token）
	gr := g.Group("", middleware.TokenVerify())
	//获取用户信息相关的路由
	userGroup := gr.Group("/user")
	{
		apiRouterGroup.UserRouter(userGroup)
	}
	//注册图片相关的路由
	imageGroup := gr.Group("/img")
	{
		apiRouterGroup.ImgRouter(imageGroup)
	}
}
