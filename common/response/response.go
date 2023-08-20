package response

import (
	"go-api/common/global"
	"github.com/gin-gonic/gin"
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
	c.JSON(http.StatusOK, Response{
		0,
		"success",
		data,
	})
}

// Fail 响应失败 ErrorCode 不为 0 表示失败
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, Response{
		code,
		msg,
		nil,
	})
}

// FailByError 失败响应 返回自定义错误的错误码、错误信息
func FailByError(c *gin.Context, error global.CustomError) {
	Fail(c, error.ErrorCode, error.ErrorMsg)
}

// ValidateFail 请求参数验证失败
func ValidateFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.ValidateError.ErrorCode, msg)
}

// BusinessFail 业务逻辑失败
func BusinessFail(c *gin.Context, msg string) {
	Fail(c, global.Errors.BusinessError.ErrorCode, msg)
}

// TokenFail Token 错误
func TokenFail(c *gin.Context) {
	FailByError(c, global.Errors.AuthError)
}

// ServerError 服务错误信息
func ServerError(c *gin.Context, err interface{}) {
	msg := "Internal Server Error"
	// 非生产环境显示具体错误信息
	if global.App.Config.App.Env != "production" && os.Getenv(gin.EnvGinMode) != gin.ReleaseMode {
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
