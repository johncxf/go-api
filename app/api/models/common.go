package models

import (
	"go-api/common/utils"
)

// ID 自增ID主键
type ID struct {
	ID uint `json:"id" gorm:"primaryKey"`
}

// Timestamps 创建更新时间
type Timestamps struct {
	//CreatedAt time.Time `json:"created_at"`
	//UpdatedAt time.Time `json:"updated_at"`
	CreatedAt utils.LocalTime `json:"created_at"`
	UpdatedAt utils.LocalTime `json:"updated_at"`
}

// SoftDeletes 软删除
type SoftDeletes struct {
	//DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
	DeletedAt utils.LocalTime `json:"deleted_at" gorm:"index"`
}
