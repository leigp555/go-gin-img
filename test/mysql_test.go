package test

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"testing"
)

type Product struct {
	gorm.Model
	Code  string
	Price uint
}

// `User` 属于 `Company`，`CompanyID` 是外键
type User struct {
	gorm.Model
	Name      string
	CompanyID int
	Company   Company `gorm:"foreignKey:CompanyID"`
}

type Company struct {
	gorm.Model
	Name string
}
type A struct {
	gorm.Model
	Name   string `gorm:"not null"`
	Age    int    `gorm:"not null"`
	Gender string `gorm:"not null"`
}

var (
	db *gorm.DB
)

func TestMysql(t *testing.T) {
	LinkDb()
	CreateTable()
	u := A{}
	ret := db.Model(A{}).Where("name=?", "hello").First(&u)
	//if ret.Error!= nil && errors.Is(gorm.ErrRecordNotFound, ret.Error) {
	//	fmt.Println("first方法没有查到数据")
	//}

	if ret.Error != nil {
		fmt.Println("find方法查询出错")
	}
	if ret.RowsAffected == 0 {
		fmt.Println("没有查到数据")
	}
	fmt.Println(u)
}

// 连接数据库
func LinkDb() {
	dsn := "root:123456abc@tcp(1.117.141.66:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	mdb, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	db = mdb
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功")
}

// 创建表
func CreateTable() {
	err := db.AutoMigrate(&Product{}, &User{}, &Company{}, &A{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("表创建成功")
}

// 创建表数据
func Create() {
	//创建单条记录
	if err := db.Create(&Product{Code: "DT55", Price: 400}).Error; err != nil {
		log.Fatal(err)
	}

	//创建多条记录
	proArr := []*Product{
		{Code: "DT60", Price: 400},
		{Code: "DT61", Price: 500},
		{Code: "DT62", Price: 600},
	}
	if err := db.Create(&proArr).Error; err != nil {
		log.Fatal(err)
	}
	//指定字段创建
	if err := db.Select("Code").Create(&Product{Code: "DT63", Price: 700}).Error; err != nil {
		log.Fatal(err)
	}

	if err := db.Omit("Code").Create(&Product{Code: "DT63", Price: 700}).Error; err != nil {
		log.Fatal(err)
	}
}

// 软删除表数据
func Delete() {
	if err := db.Where("code = ?", "D42").Delete(&Product{}).Error; err != nil {
		log.Fatal(err)
	}
}

// 永久删除表数据
func DeleteForever() {
	if err := db.Where("code = ?", "D42").Unscoped().Delete(&Product{}).Error; err != nil {
		log.Fatal(err)
	}
}

// 修改数据
func Update() {
	err := db.Model(&Product{}).Where("code=?", "D122").Updates(map[string]interface{}{"code": "DXX", "price": 456}).Error
	if err != nil {
		log.Fatal(err)
	}
}

// 查询数据
func Search() {
	var product Product
	err := db.Where("code = ?", "D44").First(&product).Error
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(product)
}
