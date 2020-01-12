package database

import "market/internal/pkg/utils"

// BaseModel 默认model结构体
type BaseModel struct {
	ID        uint           `gorm:"primary_key" json:"id"`
	CreatedAt utils.JsonTime `gorm:"type:datetime;column:create_time" json:"create_time"`
	UpdatedAt utils.JsonTime `gorm:"type:datetime;column:update_time" json:"update_time"`
}
