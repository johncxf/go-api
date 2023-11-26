package global

import (
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go-api/common/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	Logger      *zap.Logger
	ConfigViper *viper.Viper
	Config      config.Configuration
	DB          *gorm.DB
	Redis       *redis.Client
)
