package main

import (
	"fmt"
	"img/server/core"
	"img/server/global"
)

func main() {
	core.InitConf()
	fmt.Println(global.Config)
	fmt.Println("end ... ")
}
