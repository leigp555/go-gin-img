package server

import (
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/routers"
)

// StartServer 业务服务器
func (server) StartServer() {
	ginConf := global.Config.Gin
	options := options{
		Host: ginConf.Host,
		Port: ginConf.Port,
		Mode: ginConf.Mode,
		Name: "Gin",
	}
	Server.Engine(options, func(e *gin.Engine) {
		routers.InitRouter(e)
	})
}
