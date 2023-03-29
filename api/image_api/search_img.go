package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImgApi) SearchImg(c *gin.Context) {
	file := c.Query("fileName")
	imgUrl := fmt.Sprintf("http://localhost:8080/static/%s", file)
	c.JSON(200, gin.H{"msg": "success", "data": imgUrl})
}
