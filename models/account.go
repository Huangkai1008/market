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
	IsDefault   *bool  `gorm:"type:tinyint(1);index:is_default" json:"is_default"` // 是否默认地址
}

// 获取用户收货地址
func GetAddresses(maps interface{}) (addresses []Address, err error) {
	err = db.Where(maps).Find(&addresses).Error
	return addresses, err
}

// 创建用户收货地址事务
func CreateAddressTx(address Address) (Address, error) {

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if err := tx.Error; err != nil {
		return Address{}, err
	}

	// 如果设置此为默认地址, 更新其他默认地址为非默认
	if *address.IsDefault == true {
		maps := make(map[string]interface{})
		maps["is_default"] = true
		maps["user_id"] = address.UserID

		var defaultAddress Address

		if !tx.Where(maps).First(&defaultAddress).RecordNotFound() {
			tx.Model(&defaultAddress).Update(map[string]interface{}{"is_default": false})
		}

	}

	if err := tx.Create(&address).Error; err != nil {
		tx.Rollback()
		return Address{}, err
	}

	return address, tx.Commit().Error

}
