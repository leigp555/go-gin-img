package routers

import (
	"github.com/gin-gonic/gin"
	"img/server/api"
)

func (ApiRouterGroup) ArticleRouter(r *gin.RouterGroup) {

	r.GET("", func(c *gin.Context) {
		c.JSON(200, gin.H{"name": "article"})
	})
	r.GET("/title", api.GroupApi.ArticleApi.Title)
	r.GET("/content", api.GroupApi.ArticleApi.Content)
}
