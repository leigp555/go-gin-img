package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"img/server/global"
	"time"
)

type Writer struct {
}

func (w Writer) Printf(format string, args ...interface{}) {
	fmt.Println(len(args))
	if len(args) == 4 {
		global.Mlog.Info(args[3])
	} else {
		global.Mlog.Info(args)
	}

}

// LinkMysqlDB LinkDB 连接mysql数据库
func LinkMysqlDB() {
	newLogger := logger.New(
		//log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		Writer{},
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: false,       // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 禁用彩色打印
		},
	)
	// 获取mysql的配置
	var mysqlConf = global.Config.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True", mysqlConf.Username, mysqlConf.Password, mysqlConf.Addr, mysqlConf.DB)
	d, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger, //配置一个日志
	})
	if err != nil {
		global.Slog.Fatalf("mysql数据库连接失败%v\n", err)
	}
	sqlDb, _ := d.DB()
	//设置连接池
	sqlDb.SetMaxIdleConns(global.Config.Mysql.MaxConn) //设置最大连接数
	sqlDb.SetMaxOpenConns(global.Config.Mysql.MaxOpen) //设置最大的空闲连接数
	sqlDb.SetConnMaxLifetime(time.Hour * 4)            //连接最大复用时间，不能超过wait_timeout

	global.Mdb = d
	global.Slog.Info("成功连接mysql数据库")
}
