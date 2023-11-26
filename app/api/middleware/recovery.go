package middleware

import (
	"github.com/gin-gonic/gin"
	"go-api/common/global"
	"go-api/common/response"
	"gopkg.in/natefinch/lumberjack.v2"
)

// CustomRecovery 重写
func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   global.Config.Logger.RootDir + "/" + global.Config.Logger.Filename,
			MaxSize:    global.Config.Logger.MaxSize,
			MaxBackups: global.Config.Logger.MaxBackups,
			MaxAge:     global.Config.Logger.MaxAge,
			Compress:   global.Config.Logger.Compress,
		},
		response.ServerError)
}
