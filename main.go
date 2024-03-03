package main

import (
	"go-api/bootstrap"
	"go-api/common/global"
	"go-api/common/logger"
	"go.uber.org/zap"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.Logger, _ = logger.Init()

	// 初始化数据库
	global.DB = bootstrap.InitDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.DB != nil {
			db, _ := global.DB.DB()
			if err := db.Close(); err != nil {
				global.Logger.Error("close database connect failed:", zap.Error(err))
				return
			}
			global.Logger.Info("close database connect")
		}
	}()

	// 初始化验证器
	bootstrap.InitValidator()

	// 初始化 Redis
	global.Redis = bootstrap.InitRedis()

	// 启动服务器
	bootstrap.RunServer()
}
