package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (ApiRouterGroup) ImgRouter(r *gin.RouterGroup) {

	{
		r.POST("/upload", api.GroupApi.ImgApi.UploadImg) //上传图片
		r.GET("/search", api.GroupApi.ImgApi.Search)     //获取图片
	}
}
