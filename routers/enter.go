package routers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"img/server/middleware"
	"img/server/utils"
)

type ApiRouterGroup struct{}

var apiRouterGroup = new(ApiRouterGroup)

func InitRouter(r *gin.Engine) {
	//添加全局跨域中间件
	r.Use(middleware.Cors(), middleware.Log())
	//配置路由路口
	g := r.Group("v1/api")

	//测试路由
	{
		g.GET("/email", func(c *gin.Context) {
			err := utils.Email.Send([]string{"122974945@qq.com"}, "123456")
			if err != nil {
				fmt.Println("email发送失败")
				c.JSON(500, gin.H{"msg": "邮件发送失败请重试"})
			} else {
				c.JSON(200, gin.H{"msg": "邮件已发送请注意查收"})

			}
		})
		g.GET("captcha", func(c *gin.Context) {
			id, captcha, err := utils.Captcha.Generate()
			if err != nil {
				c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常，请重试"}}})
			}
			c.JSON(200, gin.H{"code": 200, "msg": "验证码获取成功", "data": map[string]any{"captchaId": id, "captchaImg": captcha}})
		})
		g.GET("captcha/verify", func(c *gin.Context) {
			b := utils.Captcha.Verify("nfXqRGroMgKrYGmDGSqA", "6360")
			if b == false {
				c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"验证码解析失败"}}})
			}
			c.JSON(200, gin.H{"code": 200, "msg": "验证码获取成功", "data": map[string]any{"msg": "验证码解析成功"}})
		})
		g.GET("md5", func(c *gin.Context) {
			s := utils.Md5Str("23456778")
			c.JSON(200, gin.H{"code": 200, "msg": "md5", "data": s})
		})
		g.GET("token", func(c *gin.Context) {
			token, err := utils.Token.Generate("lgp")
			if err != nil {
				c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常,请稍后再试"}}})
				return
			}
			c.JSON(200, gin.H{"code": 200, "msg": "md5", "token": token})
		})
		g.GET("token/verify", func(c *gin.Context) {
			t := c.Query("token")
			username, err := utils.Token.Parse(t)
			if err != nil {
				c.JSON(500, gin.H{"code": 500, "errors": map[string]any{"body": []string{"服务器异常,请稍后再试"}}})
				return
			}
			c.JSON(200, gin.H{"code": 200, "msg": "token", "username": username})
		})
	}
	//swagger路由
	{
		g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	}

	//用户登录注册路由
	type NewUserInfo struct {
		Username         string `form:"username" binding:"required,min=1,max=20" msg:"用户名不能为空,且长度为1~20位"`
		Email            string `form:"email" binding:"required,email" msg:"请输入正确的邮箱"`
		EmailCaptchaCode string `form:"emailCaptcha" binding:"required,len=6" msg:"邮箱验证码不正确"`
		Password         string `form:"password" binding:"required,min=6,max=12" msg:"密码不能为空,且长度为6~12位"`
		RePassword       string `form:"rePassword" binding:"required,min=6,max=12,eqfield=Password" msg:"两次输入的密码不一致"`
		ImgCaptchaId     string `form:"imgCaptchaId" binding:"required" msg:"图形验证码不正确"`
		ImgCaptcha       string `form:"imgCaptcha" binding:"required" msg:"图形验证码不正确"`
	}
	{
		g.POST("register", func(c *gin.Context) {
			var newUserInfo NewUserInfo
			err := c.ShouldBind(&newUserInfo)
			if err != nil {
				msg := utils.GetValidMsg(err, &newUserInfo)
				c.JSON(400, gin.H{"code": 400, "errors": map[string]any{"body": msg}})
				return
			}
			c.JSON(200, gin.H{"code": 200, "msg": "注册成功", "body": newUserInfo})
		})
	}

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
