package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (routerGroup) img(r *gin.RouterGroup) {

	{
		r.POST("/upload", api.Group.Img.UploadImg) //上传图片
		r.GET("/search", api.Group.Img.Search)     //获取图片
	}
}
