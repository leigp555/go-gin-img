package main

import (
	"img/server/core"
)

func main() {
	core.InitConf()
	core.LinkMysqlDB()
	core.LinkRedisDB()
}
