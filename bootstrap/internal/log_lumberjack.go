package internal

import (
	"go-api/common/global"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

var LogLumberjack = new(logLumberjack)

type logLumberjack struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (l *logLumberjack) GetWriteSyncer(level string) zapcore.WriteSyncer {
	file := &lumberjack.Logger{
		//Filename:   global.App.Config.Logger.RootDir + "/" + global.App.Config.Logger.Filename,
		Filename:   global.App.Config.Logger.RootDir + "/" + level + ".log",
		MaxSize:    global.App.Config.Logger.MaxSize,
		MaxBackups: global.App.Config.Logger.MaxBackups,
		MaxAge:     global.App.Config.Logger.MaxAge,
		Compress:   global.App.Config.Logger.Compress,
	}
	if global.App.Config.Logger.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	}
	return zapcore.AddSync(file)
}
