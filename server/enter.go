package server

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"img/server/core"
	"img/server/global"
	"img/server/middleware"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"time"
)

type server struct{}

var (
	wg     sync.WaitGroup
	Server = new(server)
)

func (s server) Start() {
	//初始化依赖
	core.InitDeps()
	wg.Add(2)
	//业务服务
	go func() {
		s.StartServer()
		wg.Done()
	}()
	//socket服务
	go func() {
		s.StartSocketServer()
		wg.Done()
	}()
	wg.Wait()
}

type options struct {
	Host string
	Port string
	Mode string
	Name string
}

// Engine  生成一个gin引擎处理业务路由和socket路由
func (server) Engine(options options, initRouter func(e *gin.Engine)) {
	//配置gin
	gin.SetMode(options.Mode)
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Cors())
	//初始化路由
	initRouter(r)
	//监听端口
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", options.Host, options.Port),
		Handler: r,
	}
	global.Slog.Infof("%s server 成功监听%s端口", options.Name, options.Port)
	//服务启停
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Slog.Fatalf("%s server listen error: %s\n", options.Name, err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.Slog.Infof("Shutdown %s server ...", options.Name)

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Slog.Fatalf("%s server Shutdown error:%v", options.Name, err)
	}
	global.Slog.Infof("%s server exiting", options.Name)
}
