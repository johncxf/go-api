package main

import (
	"gin-practice/bootstrap"
	"gin-practice/common/global"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.App.Logger = bootstrap.InitLogger()
	//global.App.Logger.Info("--- Service start ---")

	// 初始化数据库
	global.App.DB = bootstrap.InitDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitValidator()

	// 初始化 Redis
	global.App.Redis = bootstrap.InitRedis()

	// 启动服务器
	bootstrap.RunServer()
}
