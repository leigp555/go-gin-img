package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (routerGroup) user(r *gin.RouterGroup) {
	r.GET("/login", api.Group.User.UserLogin)
	r.GET("/register", api.Group.User.UserRegister)
}
