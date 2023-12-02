package models

import (
	"strconv"
	"time"
)

type User struct {
	ID
	Username     string     `json:"username" gorm:"not null;size:64;comment:用户名"`
	Email        string     `json:"email" gorm:"not null;size:128;comment:邮箱"`
	Mobile       string     `json:"mobile" gorm:"index;size:11;comment:手机号"`
	Password     string     `json:"password" gorm:"not null;size:128;comment:用户密码"`
	Salt         string     `json:"salt"`
	Nickname     string     `json:"nickname"`
	Avatar       string     `json:"avatar"`
	Level        int        `json:"level"`
	Sex          int        `json:"sex"`
	Birthday     *time.Time `json:"birthday"`
	Signature    string     `json:"signature"`
	UserStatus   int        `json:"user_status"`
	MobileStatus int        `json:"mobile_status"`
	UserType     int        `json:"user_type"`
	//Successions    int        `json:"successions"`
	//MaxSuccessions int        `json:"max_successions"`
	//LoginFailure   int        `json:"login_failure"`
	//LastLoginIp    string     `json:"last_login_ip"`
	//LastLoginTime  time.Time  `json:"last_login_time"`
	Timestamps
	SoftDeletes
}

func (user User) GetUid() string {
	return strconv.Itoa(int(user.ID.ID))
}
