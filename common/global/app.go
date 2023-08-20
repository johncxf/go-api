package global

import (
	"gin-practice/common/config"
	"github.com/go-redis/redis/v8"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Logger      *zap.Logger
	DB          *gorm.DB
	Redis       *redis.Client
}

var App = new(Application)
