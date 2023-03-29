package image_api

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"log"
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

	c.Next()

	filePath := fmt.Sprintf("./public/%s", file.Filename)
	src, err := imaging.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
	}
	src = imaging.Resize(src, 400, 0, imaging.Lanczos)
	img1 := imaging.Blur(src, 2)

	// 保存缩略图到文件
	err = imaging.Save(img1, "./public/thumb.jpg")
	if err != nil {
		log.Fatal(err)
	}
}
