package internal

import (
	"go-api/app/api/models"
	"go-api/common/global"
	"go.uber.org/zap"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"os"
)

// InitMysqlGorm 初始化 Mysql 数据库
func InitMysqlGorm() *gorm.DB {
	global.Logger.Info("init mysql")
	mysqlConf := global.Config.Mysql
	if mysqlConf.DBName == "" {
		global.Logger.Error("mysql config error, db_name is empty")
		return nil
	}
	// 创建 mysql.Config 实例，其中包含了连接数据库所需的信息，比如 DSN (数据源名称)，字符串类型字段的默认长度以及自动根据版本进行初始化等参数。
	mysqlConfig := mysql.Config{
		DSN:                       mysqlConf.Dsn(), // DSN data source name
		DefaultStringSize:         191,             // string 类型字段的默认长度
		DisableDatetimePrecision:  true,            // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,            // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,            // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false,           // 根据版本自动配置
	}
	// 打开数据库连接
	db, err := gorm.Open(mysql.New(mysqlConfig), Gorm.Config(mysqlConf.Prefix, mysqlConf.Singular))
	// 将引擎设置为我们配置的引擎，并设置每个连接的最大空闲数和最大连接数。
	if err != nil {
		global.Logger.Error("mysql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		db.InstanceSet("gorm:table_options", "ENGINE="+mysqlConf.Engine)
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(mysqlConf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(mysqlConf.MaxOpenConns)
		//initMySqlTables(db)
		return db
	}
}

// 数据库表初始化
func initMySqlTables(db *gorm.DB) {
	err := db.AutoMigrate(
		models.User{},
	)
	if err != nil {
		global.Logger.Error("migrate table failed", zap.Any("err", err))
		os.Exit(0)
	}
}
