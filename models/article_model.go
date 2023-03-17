package models

import (
	"gorm.io/gorm"
	"img/server/global"
)

type Article struct {
	gorm.Model
	Title   string
	Content string
}

// Generate 创建文章表
func (Article) Generate() {
	var db = global.Mydb
	err := db.AutoMigrate(&Article{})
	if err != nil {
		global.SugarLog.Panicln("Article表创建失败")
	}
}

var ArticleTable = Article{}
