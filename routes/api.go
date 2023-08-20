package routes

import (
	"gin-practice/app/api/controllers/auth"
	"gin-practice/app/api/controllers/v1"
	"gin-practice/app/api/middleware"
	"gin-practice/app/api/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// SetApiGroupRoutes 定义 api 分组路由
func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.GET("/hello", func(c *gin.Context) {
		c.String(http.StatusOK, "hello, word")
	})
	router.GET("/test", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		c.String(http.StatusOK, "success")
	})

	router.POST("/auth/login", auth.Login)
	router.POST("/auth/register", auth.Register)
	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/logout", auth.Logout)
	}
	v1AuthRouter := router.Group("/v1").Use(middleware.JWTAuth(services.AppGuardName))
	{
		v1AuthRouter.POST("/user/info", v1.Info)
	}
}
