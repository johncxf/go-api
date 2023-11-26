package bootstrap

import (
	"context"
	"github.com/go-redis/redis/v8"
	"go-api/common/global"
	"go.uber.org/zap"
)

// InitRedis 初始化 Redis 连接
func InitRedis() *redis.Client {
	global.Logger.Info("init redis")
	client := redis.NewClient(&redis.Options{
		Addr:     global.Config.Redis.Host + ":" + global.Config.Redis.Port,
		Password: global.Config.Redis.Password,
		DB:       global.Config.Redis.DB,
	})
	_, err := client.Ping(context.Background()).Result()
	if err != nil {
		global.Logger.Error("Redis connect ping failed, err:", zap.Any("err", err))
		return nil
	}
	return client
}
