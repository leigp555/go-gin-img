package models

import (
	"errors"
	"gorm.io/gorm"
	"img/server/global"
)

type UserModel struct {
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

// Generate 创建表
func (UserModel) Generate() {
	var db = global.Mdb
	err := db.AutoMigrate(&UserModel{})
	if err != nil {
		global.Slog.Panicln("User表创建失败")
	}
}

// CreateUser 创建用户
func (UserModel) CreateUser(user *UserModel) (error error) {
	if err := global.Mdb.Create(user).Error; err != nil {
		return errors.New("创建用户失败")
	}
	return nil
}

// UpdateUser 更新用户
//func (UserModel) UpdateUser() (error error) {
//	a:=map[string]any{"name":"hello","age":18,"active":false}
//	a1:=
//	global.Mdb.Model(&UserModel{}).Where().Select("name").Updates(map[string]interface{}{"name": "hello", "age": 18, "active": false}
//}

// FindUser 查询用户
//func (User) FindUser() (error error) {
//
//}
