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

func TestMysql(t *testing.T) {
	dsn := "root:123456abc@tcp(1.117.141.66:3306)/test?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("数据库连接成功")
	err = db.AutoMigrate(&Product{})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("表创建成功")
	db.Create(&Product{Code: "D42", Price: 100})
	var product Product
	db.First(&product, 1)                 // 根据整型主键查找
	db.First(&product, "code = ?", "D42") // 查找 code 字段值为 D42 的记录
	fmt.Println(product)
}
