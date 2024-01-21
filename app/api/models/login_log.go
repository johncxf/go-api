package models

import "time"

type LoginLog struct {
	ID
	UserId    uint      `json:"user_id"`
	IP        string    `json:"ip"`
	LoginTime time.Time `json:"login_time"`
}
