package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (ApiRouterGroup) UserRouter(r *gin.RouterGroup) {

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "index"})
	})
	r.GET("/login", api.GroupApi.UserApi.UserLogin)
	r.GET("/register", api.GroupApi.UserApi.UserRegister)

}
