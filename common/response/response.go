package response

import (
	"github.com/gin-gonic/gin"
	"go-api/common/global"
	"net/http"
	"os"
)

// Response 响应结构体
type Response struct {
	Code    int         `json:"code"`    // 自定义错误码
	Message string      `json:"message"` // 信息
	Data    interface{} `json:"data"`    // 数据
}

// Success 响应成功 ErrorCode 为 0 表示成功
func Success(c *gin.Context, data interface{}) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		0,
		"success",
		data,
	})
}

// Error 响应失败 ErrorCode 不为 0 表示失败
func Error(c *gin.Context, code int, msg string) {
	c.AbortWithStatusJSON(http.StatusOK, Response{
		code,
		msg,
		nil,
	})
}

// ErrorByCustomCode 失败响应 返回自定义错误的错误码、错误信息
func ErrorByCustomCode(c *gin.Context, error global.CustomError) {
	Error(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateError 请求参数验证失败
func ValidateError(c *gin.Context, msg string) {
	Error(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessError 业务逻辑失败
func BusinessError(c *gin.Context, msg string) {
	Error(c, global.Errors.BusinessError.ErrorCode, msg)
}

// TokenError Token 错误
func TokenError(c *gin.Context) {
	ErrorByCustomCode(c, global.Errors.AuthError)
}

// ServerError 服务错误信息
func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	// 非生产环境显示具体错误信息
	if global.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
		if _, ok := err.(error); ok {
			msg = err.(error).Error()
		}
	}
	c.JSON(http.StatusInternalServerError, Response{
		http.StatusInternalServerError,
		msg,
		nil,
	})
	c.Abort()
}
