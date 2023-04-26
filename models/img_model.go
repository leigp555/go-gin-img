package models

import (
	"gorm.io/gorm"
	"img/server/global"
)

type Img struct {
	gorm.Model
	ImgId     string `gorm:"type:varchar(50);comment:图片id"`
	ImgOwner  string `gorm:"type:varchar(50);comment:图片所有者"`
	ImgPath   string `gorm:"type:varchar(100);comment:图片路径"`
	ImgName   string `gorm:"type:varchar(100);comment:图片名称"`
	ThumbName string `gorm:"type:varchar(100);comment:缩略图名称"`
	ThumbPath string `gorm:"type:varchar(100);comment:缩略图路径"`
}

// Generate 创建文章表
func (Img) Generate() {
	var db = global.Mdb
	err := db.AutoMigrate(&Img{})
	if err != nil {
		global.Slog.Panicln("Img表创建失败")
	}
}

var ImgTable = Img{}
