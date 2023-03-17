package routers

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"img/server/global"
	"img/server/middleware"
	"io"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func InitRouter() {
	sysConf := global.Config.System
	//gin配置log文件
	f, _ := os.OpenFile("log/gin/log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0755)
	gin.SetMode(gin.ReleaseMode)
	gin.DefaultWriter = io.MultiWriter(f, os.Stdout)
	r := gin.Default()
	//允许跨域
	r.Use(middleware.Cors())
	g := r.Group("v1/api")
	//注册路由
	{
		g.GET("/xxx", func(c *gin.Context) {
			c.JSON(200, gin.H{"hell": "world"})
		})
	}

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
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
