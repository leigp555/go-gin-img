package main

import (
	"img/server/core"
	"img/server/global"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.LinkMysqlDB()
	core.LinkRedisDB()
	global.SugarLog.Info("xxxxxxx")
	global.Logger.Info("yyyyyyyyy")
}
