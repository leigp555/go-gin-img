package service

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/routers"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func StartServer() {
	//获取系统配置文件
	sysConf := global.Config.System
	ginConf := global.Config.Gin
	gin.SetMode(ginConf.Mode)
	//初始化gin
	//设置开发模式
	r := gin.New()
	r.Use(gin.Logger())

	//初始化路由
	routers.InitRouter(r)
	//监听端口
	srv := &http.Server{
		Addr:    fmt.Sprintf("%s:%s", sysConf.Host, sysConf.Port),
		Handler: r,
	}
	global.SugarLog.Infof("成功监听%s端口", sysConf.Port)
	//服务启停
	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.SugarLog.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.SugarLog.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.SugarLog.Fatalf("Server Shutdown:%v", err)
	}
	global.SugarLog.Info("Server exiting")
}
