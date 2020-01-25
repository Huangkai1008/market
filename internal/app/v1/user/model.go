package user

import (
	"market/internal/pkg/database/model"
	"market/internal/pkg/schema"
)

// User 用户模型
type User struct {
	model.BaseModel
	Username     string `gorm:"type:varchar(100);not null;unique;comment:'用户名'" json:"username"` // 用户名
	Email        string `gorm:"type:varchar(128);not null;unique;comment:'邮箱'" json:"email"`     // 邮箱
	HashPassword string `gorm:"type:varchar(256);not null;comment:'密码'" json:"-"`                // 密码
}

func (user *User) ToUserSchema() (readSchema *ReadSchema) {
	readSchema = &ReadSchema{
		BaseSchema: schema.BaseSchema{
			ID:        user.ID,
			CreatedAt: user.CreatedAt,
			UpdatedAt: user.UpdatedAt},
		Username: user.Username,
		Email:    user.Email,
	}
	return
}

// Address 用户收货地址模型
type Address struct {
	model.BaseModel
	UserID      uint   `gorm:"type:bigint(10);index;not null;comment:'用户id 用户1 --> 地址N'" json:"user_id"` // 用户id(用户1 --> 地址N)
	Consignee   string `gorm:"type:varchar(64);not null;comment:'收货人姓名'" json:"consignee"`               // 收货人姓名
	Mobile      string `gorm:"type:varchar(32);not null;comment:'手机号码'" json:"mobile"`                   // 手机号码
	Province    string `gorm:"type:varchar(32);not null;comment:'省份/直辖市'" json:"province"`               // 省份/直辖市
	City        string `gorm:"type:varchar(32);not null;comment:'城市'" json:"city"`                       // 市
	Region      string `gorm:"type:varchar(32);not null;comment:'所在地区'" json:"region"`                   // 所在地区
	Street      string `gorm:"type:varchar(32);not null;comment:'所在街道'" json:"street"`                   // 所在街道
	FullAddress string `gorm:"type:varchar(64);not null;comment:'详细地址'" json:"full_address"`             // 详细地址
	Tag         string `gorm:"type:varchar(32);comment:'标签'" json:"tag"`                                 // 标签
	IsDefault   *bool  `gorm:"type:tinyint(1);not null;index;comment:'是否默认地址'" json:"is_default"`        // 是否默认地址
}

type Addresses []*Address

func (address *Address) ToAddressSchema() (addressSchema *AddressSchema) {
	addressSchema = &AddressSchema{
		ID:     address.ID,
		UserID: address.UserID,
		AddressBaseSchema: AddressBaseSchema{
			Consignee:   address.Consignee,
			Mobile:      address.Mobile,
			Province:    address.Province,
			City:        address.City,
			Region:      address.Region,
			Street:      address.Street,
			FullAddress: address.FullAddress,
			Tag:         address.Tag,
			IsDefault:   address.IsDefault,
		},
	}
	return
}

func (addresses Addresses) ToAddressSchemas() []*AddressSchema {
	addressSchemas := make([]*AddressSchema, len(addresses))
	for index, address := range addresses {
		addressSchemas[index] = address.ToAddressSchema()
	}
	return addressSchemas
}
