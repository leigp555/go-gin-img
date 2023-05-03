package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (routerGroup) user(r *gin.RouterGroup) {
	//获取用户信息
	r.GET("/info", api.Group.User.UserLogin)
	//更新用户信息
	r.PUT("/info", api.Group.User.UserLogin)

	//获取用户简介
	r.GET("/profile", api.Group.User.UserRegister)
	//修改用户简介
	r.PUT("/profile", api.Group.User.UserRegister)
}
