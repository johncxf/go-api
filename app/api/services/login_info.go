package services

import (
	"go-api/app/api/models"
	"go-api/common/global"
	"time"
)

type loginLogService struct {
}

var LoginLogService = new(loginLogService)

// AddLoginLog 添加登陆日志
func (l *loginLogService) AddLoginLog(userId uint, loginIP string, loginTime time.Time) (err error, loginLog models.LoginLog) {
	loginLog = models.LoginLog{
		UserId:    userId,
		IP:        loginIP,
		LoginTime: loginTime,
	}
	err = global.DB.Create(&loginLog).Error
	return
}
