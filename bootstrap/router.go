package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-api/app/api/middleware"
	"go-api/common/global"
	"go-api/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// setupRouter 设置路由
func setupRouter() *gin.Engine {
	global.App.Logger.Info("setup router")
	//router := gin.Default()
	if global.App.Config.App.Env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if global.App.Config.App.Env == "test" {
		gin.SetMode(gin.TestMode)
	}

	router := gin.New()
	router.Use(gin.Logger(), middleware.CustomRecovery())
	// 跨域处理
	//router.Use(middleware.Cors())

	// 注册 api 分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)

	return router
}

// RunServer 启动服务器
func RunServer() {
	global.App.Logger.Info("run server")
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
	global.App.Logger.Info("--- Service stop ---")
}
