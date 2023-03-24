package models

import (
	"gorm.io/gorm"
	"img/server/global"
)

type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(20);comment:用户名"`
	Password string `gorm:"type:varchar(50);comment:密码"`
	Email    string `gorm:"type:varchar(20);unique_index;comment:邮箱"`
}

func (User) Generate() {
	var db = global.Mydb
	err := db.AutoMigrate(&User{})
	if err != nil {
		global.SugarLog.Panicln("User表创建失败")
	}
}

var UserTable = User{}
