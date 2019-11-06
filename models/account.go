package models

//用户收货地址
type Address struct {
	BaseModel
	UserID      uint   `gorm:"index:user_id;not null" json:"user_id"`              // 用户id    用户1 --> 地址N
	Consignee   string `gorm:"type:varchar(64);not null" json:"consignee"`         // 收货人姓名
	Mobile      string `gorm:"type:varchar(32);not null" json:"mobile"`            // 手机号码
	Region      string `gorm:"type:varchar(32);not null" json:"region"`            // 所在地区
	FullAddress string `gorm:"type:varchar(64);not null" json:"full_address"`      // 详细地址
	Tag         string `gorm:"type:varchar(32)" json:"tag"`                        // 标签
	IsDefault   string `gorm:"type:tinyint(1);index:is_default" json:"is_default"` // 是否默认地址
}
