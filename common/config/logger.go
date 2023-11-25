package config

import (
	"go.uber.org/zap/zapcore"
	"strings"
)

type Logger struct {
	Type          string `mapstructure:"type" json:"type" yaml:"type"`
	Level         string `mapstructure:"level" json:"level" yaml:"level"`
	RootDir       string `mapstructure:"root_dir" json:"root_dir" yaml:"root_dir"`
	Filename      string `mapstructure:"filename" json:"filename" yaml:"filename"`
	Format        string `mapstructure:"format" json:"format" yaml:"format"`
	ShowLine      bool   `mapstructure:"show_line" json:"show_line" yaml:"show_line"`
	MaxBackups    int    `mapstructure:"max_backups" json:"max_backups" yaml:"max_backups"`
	MaxSize       int    `mapstructure:"max_size" json:"max_size" yaml:"max_size"` // MB
	MaxAge        int    `mapstructure:"max_age" json:"max_age" yaml:"max_age"`    // day
	Compress      bool   `mapstructure:"compress" json:"compress" yaml:"compress"`
	Prefix        string `mapstructure:"prefix" json:"prefix" yaml:"prefix"`                         // 日志前缀
	EncodeLevel   string `mapstructure:"encode_level" json:"encode_level" yaml:"encode_level"`       // 编码级
	StacktraceKey string `mapstructure:"stacktrace_key" json:"stacktrace_key" yaml:"stacktrace_key"` // 栈名
	LogInConsole  bool   `mapstructure:"log_in_console" json:"log_in_console" yaml:"log_in_console"` // 输出控制台
}

// ZapEncodeLevel 根据 EncodeLevel 返回 zapcore.LevelEncoder
func (z *Logger) ZapEncodeLevel() zapcore.LevelEncoder {
	switch {
	case z.EncodeLevel == "LowercaseLevelEncoder": // 小写编码器(默认)
		return zapcore.LowercaseLevelEncoder
	case z.EncodeLevel == "LowercaseColorLevelEncoder": // 小写编码器带颜色
		return zapcore.LowercaseColorLevelEncoder
	case z.EncodeLevel == "CapitalLevelEncoder": // 大写编码器
		return zapcore.CapitalLevelEncoder
	case z.EncodeLevel == "CapitalColorLevelEncoder": // 大写编码器带颜色
		return zapcore.CapitalColorLevelEncoder
	default:
		return zapcore.LowercaseLevelEncoder
	}
}

// TransportLevel 将level字符串转化为 zapcore.Level
func (z *Logger) TransportLevel() zapcore.Level {
	z.Level = strings.ToLower(z.Level)
	switch z.Level {
	case "debug":
		return zapcore.DebugLevel
	case "info":
		return zapcore.InfoLevel
	case "warn":
		return zapcore.WarnLevel
	case "error":
		return zapcore.ErrorLevel
	case "dpanic":
		return zapcore.DPanicLevel
	case "panic":
		return zapcore.PanicLevel
	case "fatal":
		return zapcore.FatalLevel
	default:
		return zapcore.DebugLevel
	}
}
