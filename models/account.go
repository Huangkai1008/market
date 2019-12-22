package models

import (
	"errors"

	"market/schema"
)

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

type Addresses []*Address

func (address *Address) ToSchemaAddress() (schemaAddress *schema.Address) {
	schemaAddress = &schema.Address{
		ID:          address.ID,
		UserID:      address.UserID,
		Consignee:   address.Consignee,
		Mobile:      address.Mobile,
		Region:      address.Region,
		FullAddress: address.FullAddress,
		Tag:         address.Tag,
		IsDefault:   address.IsDefault,
	}
	return
}

func (addresses Addresses) ToSchemaAddresses() []*schema.Address {
	schemaAddresses := make([]*schema.Address, len(addresses))
	for index, address := range addresses {
		schemaAddresses[index] = address.ToSchemaAddress()
	}
	return schemaAddresses
}

// 获取用户收货地址
func GetAddresses(condition interface{}) (addresses Addresses, err error) {
	err = db.Where(condition).Find(&addresses).Error
	return addresses, err
}

// 创建用户收货地址事务
func CreateAddressTx(address *Address) (*Address, error) {

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// 如果设置此为默认地址, 更新其他默认地址为非默认
	if *address.IsDefault == true {
		condition := make(map[string]interface{})
		condition["is_default"] = true
		condition["user_id"] = address.UserID

		var defaultAddress Address

		if !tx.Where(condition).First(&defaultAddress).RecordNotFound() {
			tx.Model(&defaultAddress).Update(map[string]interface{}{"is_default": false})
		}

	}

	if err := tx.Create(&address).Error; err != nil {
		tx.Rollback()
		return &Address{}, err
	}

	return address, tx.Commit().Error

}

// 更新用户收货地址事务
func UpdateAddressTx(addressID uint, userID uint, maps map[string]interface{}) (address Address, err error) {

	tx := db.Begin()

	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	if maps["is_default"] == true {
		condition := make(map[string]interface{})
		condition["is_default"] = true
		condition["user_id"] = userID

		var defaultAddress Address

		if !tx.Where(condition).First(&defaultAddress).RecordNotFound() && defaultAddress.ID != addressID {
			tx.Model(&defaultAddress).Update(map[string]interface{}{"is_default": false})
		}
	}

	if tx.Where("id = ? and user_id = ?", addressID, userID).First(&address).RecordNotFound() {
		tx.Rollback()
		return Address{}, errors.New("不存在的收货地址")
	}

	if err := tx.Model(&address).Update(maps).Error; err != nil {
		tx.Rollback()
		return Address{}, err
	}

	return address, tx.Commit().Error

}

// 删除用户收货地址
func DeleteAddress(condition interface{}) (err error) {
	err = db.Where(condition).Delete(Address{}).Error
	return
}
