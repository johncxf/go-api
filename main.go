package main

import (
	"go-api/bootstrap"
	"go-api/common/global"
)

func main() {
	// 初始化配置
	bootstrap.InitConfig()

	// 初始化日志
	global.Logger = bootstrap.InitLogger()

	// 初始化数据库
	global.DB = bootstrap.InitDB()
	// 程序关闭前，释放数据库连接
	defer func() {
		if global.DB != nil {
			db, _ := global.DB.DB()
			db.Close()
		}
	}()

	// 初始化验证器
	bootstrap.InitValidator()

	// 初始化 Redis
	global.Redis = bootstrap.InitRedis()

	// 启动服务器
	bootstrap.RunServer()
}
