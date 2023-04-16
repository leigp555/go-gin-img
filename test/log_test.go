package test

import (
	"fmt"
	nested "github.com/antonfisher/nested-logrus-formatter"
	log "github.com/sirupsen/logrus"
	"testing"
)

func TestLog(t *testing.T) {
	//设置日志格式json||text

	// 设置json里的日期输出格式
	//log.SetFormatter(&log.JSONFormatter{
	//	TimestampFormat: "2006-01-02 15:04:05",
	//})

	// 第三方插件格式
	//type Formatter struct {
	//	FieldsOrder     []string
	//	TimestampFormat string
	//	HideKeys        bool
	//	NoColors        bool
	//	NoFieldsColors  bool
	//	ShowFullLevel   bool
	//	TrimMessages    bool
	//}

	log.SetFormatter(&nested.Formatter{
		HideKeys:        false,
		FieldsOrder:     []string{"component", "category"},
		TimestampFormat: "2006-01-02 15:04:05",
		TrimMessages:    true,
	})

	// 设置text输出格式
	//log.SetFormatter(&log.TextFormatter{
	//	TimestampFormat:           "2006-01-02 15:04:05",
	//	ForceColors:               true,
	//	EnvironmentOverrideColors: true,
	//	FullTimestamp:             true,
	//	DisableLevelTruncation:    true,
	//})

	// 设置输出警告级别
	log.SetLevel(log.DebugLevel)

	// 设置日志输出位置
	//wd, _ := os.Getwd()
	//logfile, _ := os.OpenFile(fmt.Sprintf("../%s/log/test/logrus.log",wd), os.O_CREATE|os.O_RDWR|os.O_APPEND, 0644)
	//log.SetOutput(logfile)
	//log.SetOutput(os.Stdout)

	//logrus 的 Fatal 处理
	log.RegisterExitHandler(func() {
		fmt.Println("发生了fatal异常，执行一些必要的处理工作")
	})

	//输出文件名 设置在输出日志中添加文件名和方法信息
	log.SetReportCaller(true)

	////切分日志文件
	//logger := &lumberjack.Logger{
	//	Filename:   fmt.Sprintf("%s/log/test/logrus.log", path.Join(wd, "../")),
	//	MaxSize:    500,  // 日志文件大小，单位是 MB
	//	MaxBackups: 3,    // 最大过期日志保留个数
	//	MaxAge:     28,   // 保留过期文件最大时间，单位 天
	//	Compress:   true, // 是否压缩日志，默认是不压缩。这里设置为true，压缩日志
	//}
	//log.SetOutput(logger) // logrus 设置日志的输出方式

	//示例
	log.WithFields(log.Fields{
		"animal": "walrus",
		"level":  "info",
		"xxx":    "yyy",
	}).Info("a walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"level":  "warn",
	}).Warn("a walrus appears")

	log.WithFields(log.Fields{
		"animal": "walrus",
		"level":  "error",
	}).Error("a walrus appears")
	log.WithFields(log.Fields{
		"animal": "walrus",
		"level":  "debug",
	}).Debug("a walrus appears")
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//}).Panic("a walrus appears")
	//log.WithFields(log.Fields{
	//	"animal": "walrus",
	//}).Fatal("a walrus appears")

}
