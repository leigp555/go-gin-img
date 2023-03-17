package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (ApiRouterGroup) ImgRouter(r *gin.RouterGroup) {

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "img"})
	})
	r.GET("/count", api.GroupApi.ImgApi.Count)
	r.GET("/size", api.GroupApi.ImgApi.Size)

}
