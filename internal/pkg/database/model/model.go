package model

import "market/internal/pkg/utils"

// BaseModel 默认model结构体
type BaseModel struct {
	ID        uint           `gorm:"type:bigint(10);primary_key" json:"id"`
	CreatedAt utils.JsonTime `gorm:"type:datetime;column:create_time;comment:'创建时间'" json:"create_time"`
	UpdatedAt utils.JsonTime `gorm:"type:datetime;column:update_time;comment:'更新时间'" json:"update_time"`
}
