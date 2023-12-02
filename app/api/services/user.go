package services

import (
	"errors"
	"go-api/app/api/models"
	"go-api/app/api/requests"
	"go-api/common/global"
	"go-api/common/utils"
	"strconv"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params requests.Register) (err error, user models.User) {
	emailRow := global.DB.Where("email = ?", params.Email).Select("id").First(&models.User{})
	if emailRow.RowsAffected != 0 {
		err = errors.New("该邮箱已被注册")
		return
	}
	usernameRow := global.DB.Where("username = ?", params.Username).Select("id").First(&models.User{})
	if usernameRow.RowsAffected != 0 {
		err = errors.New("该用户名已存在")
		return
	}

	password := utils.BcryptMake([]byte(params.Password))
	user = models.User{
		Username: params.Username,
		Email:    params.Email,
		Password: password}
	err = global.DB.Create(&user).Error
	return
}

// Login 登录
func (userService *userService) Login(params requests.Login) (err error, user *models.User) {
	err = global.DB.Where("email = ?", params.Email).First(&user).Error
	if err != nil || !utils.BcryptMakeCheck([]byte(params.Password), user.Password) {
		err = errors.New("账号密码错误")
	}
	return
}

// GetUserInfo 获取用户信息
func (userService *userService) GetUserInfo(id string) (err error, user models.User) {
	intId, err := strconv.Atoi(id)
	err = global.DB.First(&user, intId).Error
	//err = global.DB.Select([]string{"name", "mobile", "created_at", "updated_at"}).First(&user, intId).Error
	if err != nil {
		err = errors.New("数据不存在")
	}
	return
}
