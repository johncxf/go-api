package models

import "time"

type LoginLog struct {
	ID
	UserId    uint      `json:"userId"`
	IP        string    `json:"ip"`
	LoginTime time.Time `json:"loginTime"`
}
