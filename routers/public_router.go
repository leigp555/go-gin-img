package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"img/server/api"
)

func (ApiRouterGroup) PublicRouter(r *gin.RouterGroup) {
	//swagger文档路由
	{
		r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}
	//登录&注册
	{
		r.GET("/login", api.GroupApi.PublicApi.Login)
		r.GET("/register", api.GroupApi.PublicApi.Register)
	}

}
