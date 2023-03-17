package article_api

import "github.com/gin-gonic/gin"

func (ArticleApi) Title(c *gin.Context) {
	c.JSON(200, gin.H{"title": "pig"})
}

func (ArticleApi) Content(c *gin.Context) {
	c.JSON(200, gin.H{"count": "abc"})
}
