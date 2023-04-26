package main

import (
	"fmt"
	"github.com/panjf2000/ants/v2"
	"img/server/core"
	_ "img/server/docs"
	"img/server/global"
	"sync"
	"sync/atomic"
	"time"
)

func init() {
	//初始化配置
	//core.InitConf()
	////初始化日志
	//core.InitLogger()
	////初始化mysql
	//core.LinkMysqlDB()
	////初始化redis
	//core.LinkRedisDB()
	////生成mysql表
	//models.CreateTables()
}

// @title           Swagger Example API
// @version         2.0
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

var wg sync.WaitGroup
var counter int64
var pool *ants.Pool

func main() {
	//启动web服务
	//service.StartServer()
	core.InitConf()
	core.InitLogger()
	core.LinkMysqlDB()
	core.LinkRedisDB()
	core.LinkElasticsearch()
	core.InitPool()

	// 定义一个计数器，用于记录所有任务的总执行次数
	pool = global.Pool
	t1 := time.Now().Unix()
	// 添加100个任务到协程池中执行

	Add()
	sub()
	// 等待所有任务执行完成

	t2 := time.Now().Unix()
	wg.Wait()
	fmt.Println("time:", t2-t1)
	// 输出执行结果
	fmt.Println("counter:", counter)

}

func Add() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		pool.Submit(func() {
			// 执行任务
			atomic.AddInt64(&counter, 1)
			time.Sleep(time.Second * 1)
			defer wg.Done()
		})
	}
	fmt.Println("加法计算结束")
	//defer pool.Release()
}

func sub() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		pool.Submit(func() {
			// 执行任务
			atomic.AddInt64(&counter, -1)
			time.Sleep(time.Second * 1)
			defer wg.Done()
		})
	}
	fmt.Println("减法计算结束")
	defer pool.Release()
}
