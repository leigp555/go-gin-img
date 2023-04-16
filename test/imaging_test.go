package test

import (
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"log"
	"mime/multipart"
	"os"
	"testing"
)

func handle(c *gin.Context) {
	file, err := c.FormFile("file")
	dirname, err := os.Getwd()
	if err != nil {
		log.Fatalf("failed to get current directory: %v", err)
	}
	//打开文件
	src, err := file.Open()
	if err != nil {
		log.Fatalf("failed to open image: %v", err)
		return
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			log.Fatalf("failed to close image: %v", err)
		}
	}(src)

	// 解码上传的图像
	img, err := imaging.Decode(src)
	if err != nil {
		log.Fatalf("failed to decode image: %v", err)
		return
	}
	//调整缩略图参数
	thumb := imaging.Resize(img, 400, 0, imaging.Lanczos)
	result := imaging.Blur(thumb, 2)
	// 保存缩略图到文件
	err = imaging.Save(result, "../"+dirname+"/assets/thumb.jpg")
	if err != nil {
		log.Fatalf("failed to save image: %v", err)
	}
}

func TestImaging(t *testing.T) {
	r := gin.Default()
	r.POST("/upload", handle)

	err := r.Run(":8080")
	if err != nil {
		log.Fatalln(err)
	}
}
