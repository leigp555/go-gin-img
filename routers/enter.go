package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"img/server/middleware"
)

type ApiRouterGroup struct{}

var apiRouterGroup = new(ApiRouterGroup)

func InitRouter(r *gin.Engine) {
	//添加全局跨域中间件
	r.Use(middleware.Cors())
	//配置路由路口
	g := r.Group("v1/api")
	//swagger路由
	g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//注册用户相关的路由
	userGroup := g.Group("/user")
	{
		apiRouterGroup.UserRouter(userGroup)
	}
	//注册文章相关的路由
	articleGroup := g.Group("/article")
	{
		apiRouterGroup.ArticleRouter(articleGroup)
	}
	//注册图片相关的路由
	imageGroup := g.Group("/img")
	{
		apiRouterGroup.ImgRouter(imageGroup)
	}
}
