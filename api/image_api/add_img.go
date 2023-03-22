package image_api

import "github.com/gin-gonic/gin"

func (ImgApi) Size(c *gin.Context) {
	c.JSON(200, gin.H{"size": "1984*3465"})
}

func (ImgApi) Count(c *gin.Context) {
	c.JSON(200, gin.H{"count": "100"})
}
