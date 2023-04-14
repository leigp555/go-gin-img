package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func (ImgApi) Search(c *gin.Context) {
	imgId := c.Query("imgId")
	imgUrl := fmt.Sprintf("http://localhost:8080/static/%s", imgId)
	c.JSON(200, gin.H{"msg": "success", "data": imgUrl})
}
