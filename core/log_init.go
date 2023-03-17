package core

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"img/server/config"
	"img/server/global"
	"os"
)

var logConf config.Logger
var sysConf config.System

func InitLogger() {
	logConf = global.Config.Logger
	sysConf = global.Config.System

	writeSyncer := getLogWriter()
	encoder := getEncoder()
	level := getLogLevel()
	core := zapcore.NewCore(encoder, writeSyncer, level)

	// 开启开发模式，堆栈跟踪
	caller := zap.AddCaller()
	// 开启文件及行号
	development := zap.Development()
	// 设置初始化字段,如：添加一个服务器名称
	filed := zap.Fields(zap.String("user", logConf.Prefix))
	// 构造日志
	var logger *zap.Logger
	if logConf.ShowLine {
		logger = zap.New(core, caller, development, filed)
	} else {
		logger = zap.New(core, caller, filed)
	}
	global.Logger = logger
	global.SugarLog = logger.Sugar()
}

func getEncoder() zapcore.Encoder {
	//encoderConfig := zap.NewProductionEncoderConfig()
	//encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	//encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	//return zapcore.NewConsoleEncoder(encoderConfig)

	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     zapcore.ISO8601TimeEncoder,     // ISO8601 UTC 时间格式
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	if sysConf.Env == "dev" {
		return zapcore.NewConsoleEncoder(encoderConfig)
	} else {
		return zapcore.NewJSONEncoder(encoderConfig)
	}

}

func getLogWriter() zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   logConf.Filename,
		MaxSize:    logConf.MaxSize,
		MaxBackups: logConf.MaxBackups,
		MaxAge:     logConf.MaxAge,
		Compress:   logConf.Compress,
	}
	if global.Config.System.Env == "dev" {
		return zapcore.AddSync(os.Stderr) //开发模式下输出到控制台
	} else {
		return zapcore.AddSync(lumberJackLogger) //生产环境下输出到文件
	}
	//return zapcore.NewMultiWriteSyncer(syncFile, syncConsole)  //既输出到文件又输出到控制台
}

func getLogLevel() zapcore.Level {
	var level zapcore.Level
	switch logConf.Level {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	return level
}
