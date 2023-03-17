package core

import (
	yaml "gopkg.in/yaml.v2"
	"img/server/config"
	"img/server/global"
	"io/ioutil"
	"log"
)

//读取settings.yaml的配置

func InitConf() {
	var c = &config.Config{}
	File, err := ioutil.ReadFile("settings.yaml")
	if err != nil {
		log.Printf("读取配置文件失败 #%v", err)
	}
	err = yaml.Unmarshal(File, &c)
	if err != nil {
		log.Fatalf("配置文件解析失败: %v", err)
	}
	global.Config = *c
}
