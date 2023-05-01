package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (routerGroup) UserRouter(r *gin.RouterGroup) {
	r.GET("/login", api.GroupApi.UserApi.UserLogin)
	r.GET("/register", api.GroupApi.UserApi.UserRegister)

}
