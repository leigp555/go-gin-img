package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"img/server/api"
)

func (routerGroup) PublicRouter(r *gin.RouterGroup) {
	//swagger文档路由
	{
		r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //swagger处理函数
	}
	//验证码发送
	{
		r.GET("/imgCaptcha", api.GroupApi.PublicApi.SendImgCaptcha)      //图形验证码
		r.POST("/emailCaptcha", api.GroupApi.PublicApi.SendEmailCaptcha) //邮箱验证码
	}
	//登录&注册
	{
		r.POST("/register", api.GroupApi.PublicApi.Register)      //注册
		r.GET("/login", api.GroupApi.PublicApi.Login)             //普通登录
		r.GET("/auth/github", api.GroupApi.PublicApi.GithubLogin) //GitHub登录
		r.GET("/auth/google", api.GroupApi.PublicApi.GoogleLogin) //google登录
	}
}
