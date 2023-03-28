package image_api

import (
	"github.com/gin-gonic/gin"
)

func (ImgApi) UploadImg(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"msg": "上传失败"})
		return
	}
	dst := "./public/" + file.Filename
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(500, gin.H{"msg": "存储失败"})
		return
	}
	c.JSON(200, gin.H{"file": file.Filename, "size": file.Size})
}
