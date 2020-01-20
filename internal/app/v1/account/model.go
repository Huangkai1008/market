package account

import (
	"market/internal/pkg/database/model"
)

// Address 用户收货地址模型
type Address struct {
	model.BaseModel
	UserID      uint   `gorm:"type:bigint(10);index;not null;comment:'用户id 用户1 --> 地址N'" json:"user_id"` // 用户id(用户1 --> 地址N)
	Consignee   string `gorm:"type:varchar(64);not null;comment:'收货人姓名'" json:"consignee"`               // 收货人姓名
	Mobile      string `gorm:"type:varchar(32);not null;comment:'手机号码'" json:"mobile"`                   // 手机号码
	Region      string `gorm:"type:varchar(32);not null;comment:'所在地区'" json:"region"`                   // 所在地区
	FullAddress string `gorm:"type:varchar(64);not null;comment:'详细地址'" json:"full_address"`             // 详细地址
	Tag         string `gorm:"type:varchar(32);comment:'标签'" json:"tag"`                                 // 标签
	IsDefault   *bool  `gorm:"type:tinyint(1);index;comment:'是否默认地址'" json:"is_default"`                 // 是否默认地址
}

type Addresses []*Address

func (address *Address) ToSchemaAddress() (schemaAddress *AddressSchema) {
	schemaAddress = &AddressSchema{
		ID:     address.ID,
		UserID: address.UserID,
		AddressBaseSchema: AddressBaseSchema{
			Consignee:   address.Consignee,
			Mobile:      address.Mobile,
			Region:      address.Region,
			FullAddress: address.FullAddress,
			Tag:         address.Tag,
			IsDefault:   address.IsDefault,
		},
	}
	return
}

func (addresses Addresses) ToSchemaAddresses() []*AddressSchema {
	schemaAddresses := make([]*AddressSchema, len(addresses))
	for index, address := range addresses {
		schemaAddresses[index] = address.ToSchemaAddress()
	}
	return schemaAddresses
}
