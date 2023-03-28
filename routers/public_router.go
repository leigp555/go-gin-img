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
	//验证码发送
	{
		r.GET("/imgCaptcha", api.GroupApi.PublicApi.SendImgCaptcha)
		r.POST("/emailCaptcha", api.GroupApi.PublicApi.SendEmailCaptcha)
	}
	//登录&注册
	{
		r.POST("/register", api.GroupApi.PublicApi.Register)
		r.GET("/login", api.GroupApi.PublicApi.Login)
		r.GET("/auth/github", api.GroupApi.PublicApi.GithubLogin)
		r.GET("/auth/google", api.GroupApi.PublicApi.GoogleLogin)
	}
	{
		r.POST("/upload", api.GroupApi.ImgApi.UploadImg)
		r.GET("/searchImg", api.GroupApi.ImgApi.SearchImg)
	}
}
