package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"img/server/global"
	"log"
)

// 获取mysql的配置
var mysqlConf = global.Config.Mysql

// LinkMysqlDB LinkDB 连接mysql数据库
func LinkMysqlDB() {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", mysqlConf.Username, mysqlConf.Password, mysqlConf.Addr, mysqlConf.DB)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置一个日志
	})
	if err != nil {
		log.Panicf("mysql数据库连接失败%v\n", err)
	}
	sqlDb, _ := d.DB()
	//设置连接池
	sqlDb.SetMaxIdleConns(global.Config.Mysql.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpen) //设置最大的空闲连接数

	global.Mydb = d
	fmt.Println("成功连接mysql数据库")
}
