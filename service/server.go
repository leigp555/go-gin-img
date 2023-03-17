package service

import "img/server/routers"

func StartServer() {
	//初始化路由,创建服务
	routers.InitRouter()
}
