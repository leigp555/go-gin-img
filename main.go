package main

import (
	"img/server/core"
)

func main() {
	core.InitConf()
	core.InitLogger()
	core.LinkMysqlDB()
	core.LinkRedisDB()
}
