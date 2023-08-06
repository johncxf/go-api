package routes

import (
	"gin-practice/app/controllers/api/v1"
	request2 "gin-practice/app/requests"
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
}

// SetApiV1GroupRoutes 定义 api 分组路由
func SetApiV1GroupRoutes(router *gin.RouterGroup) {
	router.POST("/user/register", func(c *gin.Context) {
		var form request2.Register
		if err := c.ShouldBindJSON(&form); err != nil {
			c.JSON(http.StatusOK, gin.H{
				"error": request2.GetErrorMsg(form, err),
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
	router.POST("/auth/register", v1.Register)
}
