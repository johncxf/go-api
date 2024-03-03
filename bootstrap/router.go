package bootstrap

import (
	"context"
	"github.com/gin-gonic/gin"
	"go-api/app/api/middleware"
	"go-api/common/global"
	"go-api/routes"
	"go.uber.org/zap"
	"net/http"
	"os"
	"os/signal"
	"time"
)

// setupRouter 设置路由
func setupRouter() *gin.Engine {
	global.Logger.Info("setup router")
	env := global.Config.App.Env
	if env == "production" {
		gin.SetMode(gin.ReleaseMode)
	} else if env == "test" {
		gin.SetMode(gin.TestMode)
	} else {
		gin.SetMode(gin.DebugMode)
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
	global.Logger.Info("run server")
	r := setupRouter()

	srv := &http.Server{
		Addr:    ":" + global.Config.App.Port,
		Handler: r,
	}

	go func() {
		// 服务连接
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			global.Logger.Error("listen error:", zap.Error(err))
		}
	}()

	// 等待中断信号以优雅地关闭服务器（设置 5 秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	global.Logger.Info("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		global.Logger.Error("Server shutdown error:", zap.Error(err))
	}
	global.Logger.Info("--- Service exiting ---")
}
