package image_api

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"log"
	"mime/multipart"
	"strings"
	"time"
)

func (ImgApi) UploadImg(c *gin.Context) {
	mdb := global.Mydb
	//获取上传的图片
	file, err := c.FormFile("file")
	if err != nil {
		c.JSON(500, gin.H{"msg": "上传失败"})
		return
	}
	//检查文件类型
	nameArr := strings.Split(file.Filename, ".")
	if len(nameArr) <= 1 || len(nameArr) > 2 {
		c.JSON(400, gin.H{"msg": "请使用正确的图片扩展名"})
		return
	}
	if nameArr[1] != "jpg" && nameArr[1] != "png" && nameArr[1] != "jpeg" && nameArr[1] != "gif" {
		c.JSON(400, gin.H{"msg": "只接收扩展名为 image/png, image/jpeg, image/gif, image/jpg的图片"})
		return
	}
	//生成图片名和图片ID以及存储路径
	username := c.GetString("userId")
	fileId := utils.Md5Str(fmt.Sprintf("%s-%s-%s", time.Now().Format("2006/01/02"), username, file.Filename))
	fileName := fileId + "." + nameArr[1]
	dst := "./public/img/" + fileName
	//缩略图名字
	thumbName := fmt.Sprintf("thumb%s", fileName)
	//生成缩略图存储路径
	thumbFilePath := fmt.Sprintf("./public/thumb/%s", thumbName)

	//检查数据库中是否已经存在该图片
	var dbImg = models.Img{}
	err = mdb.Where("img_owner = ? AND img_id=?", username, fileId).First(&dbImg).Error
	if err == nil {
		c.JSON(400, gin.H{"msg": "文件已存在"})
		return
	}
	//将图片路径以及id添加到数据库中
	newImg := models.Img{
		ImgId:     fileId,
		ImgOwner:  username,
		ImgPath:   dst,
		ImgName:   fileName,
		ThumbName: thumbName,
		ThumbPath: thumbFilePath,
	}
	if err = mdb.Create(&newImg).Error; err != nil {
		c.JSON(500, gin.H{"msg": "mysql存储失败"})
		return
	}
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		c.JSON(500, gin.H{"msg": "存储失败"})
		return
	}
	c.JSON(200, gin.H{"file": fileName, "size": file.Size})

	//响应结束生成图片缩略图
	c.Next()
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		panic("failed to open image: %v")
		return
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			panic("failed to close image: %v")
		}
	}(src)

	// 解码上传的图像
	img, err := imaging.Decode(src)

	if err != nil {
		panic("failed to open image: %v")
		return
	}
	//调整缩略图参数
	thumb := imaging.Resize(img, 400, 0, imaging.Lanczos)
	result := imaging.Blur(thumb, 2)
	// 保存缩略图到文件
	err = imaging.Save(result, thumbFilePath)
	if err != nil {
		log.Fatal(err)
	}
}
