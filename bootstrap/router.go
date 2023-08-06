package bootstrap

import (
	"context"
	"gin-practice/global"
	"gin-practice/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// setupRouter 设置路由
func setupRouter() *gin.Engine {
	router := gin.Default()

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	// 注册 api v1 分组路由
	apiV1Group := router.Group("/api/v1")
	routes.SetApiV1GroupRoutes(apiV1Group)

	return router
}

// RunServer 启动服务器
func RunServer() {
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}

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
