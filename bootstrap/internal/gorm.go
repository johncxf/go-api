package internal

import (
	"fmt"
	"go-api/common/global"
	"log"
	"os"
	"time"

	"gorm.io/gorm/schema"

	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DBBase interface {
	GetLogMode() string
}

var Gorm = new(_gorm)

type _gorm struct{}

type CustomWriter struct {
	logger.Writer
}

// Config gorm 自定义配置
func (g *_gorm) Config(prefix string, singular bool) *gorm.Config {
	// 将传入的字符串前缀和单复数形式参数应用到 GORM 的命名策略中，并禁用迁移过程中的外键约束，返回最终生成的 GORM 配置信息。
	config := &gorm.Config{
		// 命名策略
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   prefix,   // 表前缀，在表名前添加前缀，如添加用户模块的表前缀 user_
			SingularTable: singular, // 是否使用单数形式的表名，如果设置为 true，那么 User 模型会对应 users 表
		},
		// 是否在迁移时禁用外键约束，默认为 false，表示会根据模型之间的关联自动生成外键约束语句
		DisableForeignKeyConstraintWhenMigrating: true,
		Logger:                                   getGormLogger(), // 使用自定义 Logger
	}

	return config
}

// 获取 gorm 日志配置
func getGormLogger() logger.Interface {
	dbBase := getDbBase()

	var logMode logger.LogLevel
	switch dbBase.GetLogMode() {
	case "silent":
		logMode = logger.Silent
	case "error":
		logMode = logger.Error
	case "warn":
		logMode = logger.Warn
	case "info":
		logMode = logger.Info
	default:
		logMode = logger.Info
	}

	return logger.New(getGormLogWriter(), logger.Config{
		SlowThreshold:             200 * time.Millisecond, // 慢 SQL 阈值
		LogLevel:                  logMode,                // 日志级别，只记录级别不低于该值的日志
		IgnoreRecordNotFoundError: false,                  // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,                   // 禁用彩色打印
	})
}

// 自定义 gorm Writer
func getGormLogWriter() logger.Writer {
	//var writer io.Writer
	//
	//// 是否启用日志文件
	//if global.Config.Database.EnableFileLogWriter {
	//	// 自定义 Writer
	//	writer = &lumberjack.Logger{
	//		Filename:   global.Config.Logger.RootDir + "/" + global.Config.Database.LogFilename,
	//		MaxSize:    global.Config.Logger.MaxSize,
	//		MaxBackups: global.Config.Logger.MaxBackups,
	//		MaxAge:     global.Config.Logger.MaxAge,
	//		Compress:   global.Config.Logger.Compress,
	//	}
	//} else {
	//	// 默认 Writer
	//	writer = os.Stdout
	//}
	//return log.New(writer, "\r\n", log.LstdFlags)
	return newWriter(log.New(os.Stdout, "\r\n", log.LstdFlags))
}

// 获取数据库配置实例
func getDbBase() DBBase {
	var dbBase DBBase
	switch global.Config.App.DBDriver {
	case "mysql":
		dbBase = &global.Config.Mysql
	case "pgsql":
		dbBase = &global.Config.Pgsql
	default:
		dbBase = &global.Config.Mysql
	}

	return dbBase
}

// NewWriter writer 构造函数
func newWriter(w logger.Writer) *CustomWriter {
	return &CustomWriter{Writer: w}
}

// Printf 格式化打印日志
func (w *CustomWriter) Printf(message string, data ...interface{}) {
	var logZap bool
	switch global.Config.App.DBDriver {
	case "mysql":
		logZap = global.Config.Mysql.EnableFileLogWriter
	case "pgsql":
		logZap = global.Config.Pgsql.EnableFileLogWriter
	default:
		logZap = true
	}
	if logZap {
		global.Logger.Info(fmt.Sprintf(message+"\n", data...))
	} else {
		w.Writer.Printf(message, data...)
	}
}
