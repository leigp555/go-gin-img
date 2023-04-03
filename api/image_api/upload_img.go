package image_api

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/models"
	"img/server/utils"
	"mime/multipart"
	"strings"
	"time"
)

func (ImgApi) UploadImg(c *gin.Context) {
	mdb := global.Mydb
	res := utils.Res
	//获取上传的图片
	file, err := c.FormFile("file")
	if err != nil {
		res.Fail.Normal(c, 400, "请上传文件")
		return
	}
	//检查文件类型
	nameArr := strings.Split(file.Filename, ".")
	if len(nameArr) <= 1 || len(nameArr) > 2 {
		res.Fail.Normal(c, 400, "请使用正确的图片扩展名")
		return
	}
	if nameArr[1] != "jpg" && nameArr[1] != "png" && nameArr[1] != "jpeg" && nameArr[1] != "gif" {
		res.Fail.Normal(c, 400, "只接收扩展名为 image/png, image/jpeg, image/gif, image/jpg的图片")
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
		res.Fail.Normal(c, 400, "图片已存在")
		return
	}
	// 上传文件至指定的完整文件路径
	err = c.SaveUploadedFile(file, dst)
	if err != nil {
		res.Fail.Normal(c, 500, "图片存储失败")
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
		res.Fail.Normal(c, 500, "添加图片失败")
		return
	}

	//存储成功
	res.Success.Normal(c, "上传成功", map[string]string{"imgId": newImg.ImgId})

	//响应结束生成图片缩略图
	c.Next()
	// 打开上传的文件
	src, err := file.Open()
	if err != nil {
		global.SugarLog.Error("failed to open image: %v", err)
		return
	}
	defer func(src multipart.File) {
		err := src.Close()
		if err != nil {
			global.SugarLog.Error("failed to close image: %v", err)
		}
	}(src)

	// 解码上传的图像
	img, err := imaging.Decode(src)
	if err != nil {
		global.SugarLog.Error("failed to decode image: %v", err)
		return
	}
	//调整缩略图参数
	thumb := imaging.Resize(img, 400, 0, imaging.Lanczos)
	result := imaging.Blur(thumb, 2)
	// 保存缩略图到文件
	err = imaging.Save(result, thumbFilePath)
	if err != nil {
		global.SugarLog.Error("failed to save image: %v", err)
	}
}
