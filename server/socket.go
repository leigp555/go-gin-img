package server

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/routers"
)

// StartSocketServer socket服务
func (server) StartSocketServer() {
	socketConf := global.Config.Socket
	options := options{
		Host: socketConf.Host,
		Port: socketConf.Port,
		Mode: socketConf.Mode,
		Name: "Socket",
	}
	Server.Engine(options, func(e *gin.Engine) {
		//初始化路由
		routers.InitSocketRouter(e)
	})
}
