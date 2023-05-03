package models

import (
	"gorm.io/gorm"
	"img/server/global"
)

type User struct {
	gorm.Model
	Username       string `gorm:"type:varchar(50);comment:用户名"`
	NickName       string `gorm:"type:varchar(50);comment:昵称"`
	Password       string `gorm:"type:varchar(50);comment:密码"`
	Uid            string `gorm:"type:varchar(50);comment:uid"`
	Email          string `gorm:"type:varchar(50);unique_index;comment:邮箱"`
	DiskLimit      int    `gorm:"type:int;default:5120;comment:磁盘大小"`
	DiskUsage      int    `gorm:"type:int;default:0;comment:磁盘已使用"`
	Role           string `gorm:"type:varchar(50);default:normal;comment:角色"`
	EmailVerified  bool   `gorm:"type:bool;default:false;comment:邮箱验证"`
	LastLogin      string `gorm:"type:varchar(50);comment:最后登录时间"`
	LastLoginIp    string `gorm:"type:varchar(50);comment:最后登录ip"`
	CurrentLogin   string `gorm:"type:varchar(50);comment:当前登录时间"`
	CurrentLoginIp string `gorm:"type:varchar(50);comment:当前登录ip"`
	Bio            string `gorm:"type:varchar(100);comment:个人简介"`
	Avatar         string `gorm:"type:varchar(100);comment:用户头像"`
}

func (User) Generate() {
	var db = global.Mdb
	err := db.AutoMigrate(&User{})
	if err != nil {
		global.Slog.Panicln("User表创建失败")
	}
}

var UserTable = User{}
