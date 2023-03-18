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
	r.Use(middleware.Cors())
	//配置路由路口
	g := r.Group("v1/api")

	//swagger路由
	{
		g.GET("swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
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
