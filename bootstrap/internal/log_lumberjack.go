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
		//Filename:   global.Config.Logger.RootDir + "/" + global.Config.Logger.Filename,
		Filename:   global.Config.Logger.RootDir + "/" + level + ".log",
		MaxSize:    global.Config.Logger.MaxSize,
		MaxBackups: global.Config.Logger.MaxBackups,
		MaxAge:     global.Config.Logger.MaxAge,
		Compress:   global.Config.Logger.Compress,
	}
	if global.Config.Logger.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(file))
	}
	return zapcore.AddSync(file)
}
