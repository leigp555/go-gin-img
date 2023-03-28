package image_api

import (
	"github.com/gin-gonic/gin"
	"os"
)

func (ImgApi) SearchImg(c *gin.Context) {
	fileName := c.Query("fileName")
	file, err := os.ReadFile("./public/" + fileName + ".jpg")
	if err != nil {
		c.JSON(500, gin.H{"msg": "读取文件失败"})
		return
	}
	c.JSON(200, gin.H{"msg": "success", "data": file})
}
