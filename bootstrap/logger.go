package bootstrap

import (
	"go-api/bootstrap/internal"
	"go-api/common/global"
	"go-api/common/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

func InitLogger() *zap.Logger {
	// 创建根目录
	if ok, _ := utils.PathExists(global.App.Config.Logger.RootDir); !ok {
		_ = os.Mkdir(global.App.Config.Logger.RootDir, os.ModePerm)
	}

	cores := internal.Zap.GetZapCores()
	logger := zap.New(zapcore.NewTee(cores...))

	if global.App.Config.Logger.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}

	return logger
}
