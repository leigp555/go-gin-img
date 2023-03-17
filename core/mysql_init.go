package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"img/server/global"
)

// LinkMysqlDB LinkDB 连接mysql数据库
func LinkMysqlDB() {
	// 获取mysql的配置
	var mysqlConf = global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", mysqlConf.Username, mysqlConf.Password, mysqlConf.Addr, mysqlConf.DB)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), //配置一个日志
	})
	if err != nil {
		global.SugarLog.Fatalf("mysql数据库连接失败%v\n", err)
	}
	sqlDb, _ := d.DB()
	//设置连接池
	sqlDb.SetMaxIdleConns(global.Config.Mysql.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpen) //设置最大的空闲连接数

	global.Mydb = d
	global.SugarLog.Info("成功连接mysql数据库")
}
