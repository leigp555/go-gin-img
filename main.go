package main

import (
	"img/server/core"
	_ "img/server/docs"
	"img/server/models"
	"img/server/service"
)

// @title           Swagger Example API
// @version         1.0
// @description     This is a sample server celler server.
// @termsOfService  http://swagger.io/terms/

// @contact.name   API Support
// @contact.url    http://www.swagger.io/support
// @contact.email  support@swagger.io

// @license.name  Apache 2.0
// @license.url   http://www.apache.org/licenses/LICENSE-2.0.html

// @host      localhost:8080
// @BasePath  /api/v1

// @securityDefinitions.basic  BasicAuth

// @externalDocs.description  OpenAPI
// @externalDocs.url          https://swagger.io/resources/open-api/

func main() {
	//初始化配置
	core.InitConf()
	//初始化日志
	core.InitLogger()
	//初始化mysql
	core.LinkMysqlDB()
	//初始化redis
	core.LinkRedisDB()
	//生成mysql表
	models.CreateTables()
	//启动web服务
	service.StartServer()

}
