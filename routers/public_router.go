package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"img/server/api"
)

func (routerGroup) public(r *gin.RouterGroup) {
	//swagger文档路由
	{
		r.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler)) //swagger处理函数
	}
	//验证码发送
	{
		r.GET("/captcha", api.Group.Public.SendImgCaptcha)  //图形验证码
		r.POST("/email", api.Group.Public.SendEmailCaptcha) //邮箱验证码
	}
	//登录&注册
	{
		r.POST("/register", api.Group.Public.Register)      //注册
		r.GET("/login", api.Group.Public.Login)             //普通登录
		r.GET("/auth/github", api.Group.Public.GithubLogin) //GitHub登录
		r.GET("/auth/google", api.Group.Public.GoogleLogin) //google登录
	}
}
