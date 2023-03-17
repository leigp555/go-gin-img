package main

import (
	"img/server/core"
	"img/server/service"
)

func main() {
	//初始化配置
	core.InitConf()
	//初始化日志
	core.InitLogger()
	//初始化mysql
	core.LinkMysqlDB()
	//初始化redis
	core.LinkRedisDB()
	//启动web服务
	service.StartServer()

}
