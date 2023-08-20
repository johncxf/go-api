package bootstrap

import (
	"context"
	"gin-practice/common/global"
	"github.com/go-redis/redis/v8"
	"go.uber.org/zap"
)

// InitRedis 初始化 Redis 连接
func InitRedis() *redis.Client {
	client := redis.NewClient(&redis.Options{
		Addr:     global.App.Config.Redis.Host + ":" + global.App.Config.Redis.Port,
		Password: global.App.Config.Redis.Password,
		DB:       global.App.Config.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.App.Logger.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
