package middleware

import (
	"gin-practice/common/global"
	"gin-practice/common/response"
	"github.com/gin-gonic/gin"
	"gopkg.in/natefinch/lumberjack.v2"
)

// CustomRecovery 重写
func CustomRecovery() gin.HandlerFunc {
	return gin.RecoveryWithWriter(
		&lumberjack.Logger{
			Filename:   global.App.Config.Logger.RootDir + "/" + global.App.Config.Logger.Filename,
			MaxSize:    global.App.Config.Logger.MaxSize,
			MaxBackups: global.App.Config.Logger.MaxBackups,
			MaxAge:     global.App.Config.Logger.MaxAge,
			Compress:   global.App.Config.Logger.Compress,
		},
		response.ServerError)
}
