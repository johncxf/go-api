package bootstrap

import (
	"go-api/bootstrap/internal"
	"go-api/common/global"
	"gorm.io/gorm"
)

func InitDB() *gorm.DB {
	// 根据驱动配置进行初始化
	switch global.Config.App.DBDriver {
	case "mysql":
		return internal.InitMysqlGorm()
	case "pgsql":
		return internal.InitPgSqlGorm()
	default:
		return internal.InitMysqlGorm()
	}
}
