package services

import (
	"errors"
	"gin-practice/app/models"
	"gin-practice/app/requests"
	"gin-practice/common"
	"gin-practice/global"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params requests.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号已存在")
		return
	}
	user = models.User{Name: params.Name, Mobile: params.Mobile, Password: common.BcryptMake([]byte(params.Password))}
	err = global.App.DB.Create(&user).Error
	return
}
