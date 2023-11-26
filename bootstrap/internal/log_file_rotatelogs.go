package internal

import (
	rotatelogs "github.com/lestrrat-go/file-rotatelogs"
	"go-api/common/global"
	"go.uber.org/zap/zapcore"
	"os"
	"path"
	"time"
)

var LogFileRotateLogs = new(logFileRotateLogs)

type logFileRotateLogs struct{}

// GetWriteSyncer 获取 zapcore.WriteSyncer
func (r *logFileRotateLogs) GetWriteSyncer(level string) (zapcore.WriteSyncer, error) {
	fileWriter, err := rotatelogs.New(
		path.Join(global.Config.Logger.RootDir, "%Y-%m-%d", level+".log"),
		rotatelogs.WithClock(rotatelogs.Local),
		rotatelogs.WithMaxAge(time.Duration(global.Config.Logger.MaxAge)*24*time.Hour), // 日志留存时间
		rotatelogs.WithRotationTime(time.Hour*24),
	)
	if global.Config.Logger.LogInConsole {
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(fileWriter)), err
	}
	return zapcore.AddSync(fileWriter), err
}
