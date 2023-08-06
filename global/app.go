package global

import (
	"gin-practice/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Logger      *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
