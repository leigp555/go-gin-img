package routers

import (
	"github.com/gin-gonic/gin"
)

func (ApiRouterGroup) ImgRouter(r *gin.RouterGroup) {

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "img"})
	})
}
