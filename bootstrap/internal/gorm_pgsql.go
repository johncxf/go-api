package internal

import (
	"go-api/common/global"
	"go.uber.org/zap"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// InitPgSqlGorm 初始化 Postgresql 数据库
func InitPgSqlGorm() *gorm.DB {
	global.Logger.Info("init pgsql")
	pgsqlConf := global.Config.Pgsql
	if pgsqlConf.DBName == "" {
		global.Logger.Error("pgsql config error, db_name is empty")
		return nil
	}
	pgsqlConfig := postgres.Config{
		DSN:                  pgsqlConf.Dsn(), // DSN data source name
		PreferSimpleProtocol: false,
	}
	db, err := gorm.Open(postgres.New(pgsqlConfig), Gorm.Config(pgsqlConf.Prefix, pgsqlConf.Singular))
	if err != nil {
		global.Logger.Error("pgsql connect failed, err:", zap.Any("err", err))
		return nil
	} else {
		sqlDB, _ := db.DB()
		sqlDB.SetMaxIdleConns(pgsqlConf.MaxIdleConns)
		sqlDB.SetMaxOpenConns(pgsqlConf.MaxOpenConns)
		return db
	}
}
