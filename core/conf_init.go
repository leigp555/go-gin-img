package core

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"img/server/config"
	"img/server/global"
	"log"
)

//读取settings.yaml的配置

func InitConf() {
	var c = &config.Config{}
	//读取配置文件
	viper.SetConfigFile("config/settings.yaml")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("配置文件读取失败, err: %v", err)
	}
	// 反序列化配置
	if err := viper.Unmarshal(c); err != nil {
		panic(fmt.Errorf("配置文件解析失败, err:%s \n", err))
	}
	// 监控配置文件变化
	viper.WatchConfig()
	// 注意！！！配置文件发生变化后要同步到全局变量Conf
	viper.OnConfigChange(func(in fsnotify.Event) {
		if err := viper.Unmarshal(c); err != nil {
			panic(fmt.Errorf("配置文件解析失败, err:%s \n", err))
		}
		log.Println("已修改配置文件")
	})
	global.Config = c
}
