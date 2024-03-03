package logger

import (
	"go-api/common/global"
	"go-api/common/utils"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
)

var (
	Logger      *zap.Logger
	SugarLogger *zap.SugaredLogger
)

func Init() (*zap.Logger, *zap.SugaredLogger) {
	// 创建根目录
	if ok, _ := utils.PathExists(global.Config.Logger.RootDir); !ok {
		_ = os.Mkdir(global.Config.Logger.RootDir, os.ModePerm)
	}

	cores := Zap.GetZapCores()
	Logger = zap.New(zapcore.NewTee(cores...))
	if global.Config.Logger.ShowLine {
		Logger = Logger.WithOptions(zap.AddCaller())
	}
	SugarLogger = Logger.Sugar()
	return Logger, SugarLogger
}
