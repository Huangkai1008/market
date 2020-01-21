package user

import (
	"fmt"

	"gopkg.in/go-playground/validator.v9"

	"market/internal/pkg/ecode"
	"market/internal/pkg/schema"
)

type RegisterSchema struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=256"`
}

type LoginSchema struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *RegisterSchema) Validate(errs validator.ValidationErrors) ecode.MarketError {
	var marketError ecode.MarketError
	err := errs[0] // 获得第一个错误并返回, yield错误

	if err.Field() == "Username" {
		switch err.Tag() {
		case "required":
			marketError.Message = "用户名不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("用户名长度不能超过%s个字符", err.Param())
		}

	} else if err.Field() == "Password" {
		switch err.Tag() {
		case "required":
			marketError.Message = "密码不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("密码长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Email" {
		switch err.Tag() {
		case "required":
			marketError.Message = "邮箱不能为空"
		case "email":
			marketError.Message = "邮箱格式不正确"
		case "max":
			marketError.Message = fmt.Sprintf("邮箱长度不能超过%s个字符", err.Param())
		}
	}

	return marketError
}

func (l *LoginSchema) Validate(errs validator.ValidationErrors) ecode.MarketError {
	var marketError ecode.MarketError
	err := errs[0] // 获得第一个错误并返回, yield错误
	if err.Field() == "Username" {
		switch err.Tag() {
		case "required":
			marketError.Message = "用户名不能为空"
		}

	} else if err.Field() == "Password" {
		switch err.Tag() {
		case "required":
			marketError.Message = "密码不能为空"
		}
	}

	return marketError
}

type ReadSchema struct {
	schema.BaseSchema
	Username string `json:"username"`
	Email    string `json:"email"`
}

type TokenBackSchema struct {
	Token string `json:"token"`
}

type AddressBaseSchema struct {
	Consignee   string `json:"consignee" validate:"required,max=64"`
	Mobile      string `json:"mobile" validate:"required,max=32"`
	Province    string `json:"province" validate:"required,max=32"`
	City        string `json:"city" validate:"required,max=32"`
	Region      string `json:"region" validate:"required,max=32"`
	Street      string `json:"street" validate:"required,max=32"`
	FullAddress string `json:"full_address" validate:"required,max=64"`
	Tag         string `json:"tag" validate:"max=32"`
	IsDefault   *bool  `json:"is_default" validate:"required"`
}
type (
	// AddressCreateSchema 收货地址创建schema
	AddressCreateSchema struct {
		AddressBaseSchema
	}

	// AddressUpdateSchema 收货地址更新schema
	AddressUpdateSchema struct {
		AddressBaseSchema
	}
)

func (a *AddressBaseSchema) ToAddress() (address *Address) {
	address = &Address{
		Consignee:   a.Consignee,
		Mobile:      a.Mobile,
		Province:    a.Province,
		City:        a.City,
		Region:      a.Region,
		Street:      a.Street,
		FullAddress: a.FullAddress,
		Tag:         a.Tag,
		IsDefault:   a.IsDefault,
	}
	return
}

func (a *AddressBaseSchema) ToMap() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["consignee"] = a.Consignee
	maps["mobile"] = a.Mobile
	maps["province"] = a.Province
	maps["city"] = a.City
	maps["region"] = a.Region
	maps["street"] = a.Street
	maps["full_address"] = a.Consignee
	maps["tag"] = a.Tag
	maps["is_default"] = *a.IsDefault
	return maps
}

func (a *AddressBaseSchema) Validate(errs validator.ValidationErrors) ecode.MarketError {
	var marketError ecode.MarketError
	err := errs[0]
	if err.Field() == "Consignee" {
		switch err.Tag() {
		case "required":
			marketError.Message = "收货人姓名不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("收货人姓名长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Mobile" {
		switch err.Tag() {
		case "required":
			marketError.Message = "收件人手机号码不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("收件人手机号码长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Province" {
		switch err.Tag() {
		case "required":
			marketError.Message = "省份/直辖市不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("省份/直辖市字段长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "City" {
		switch err.Tag() {
		case "required":
			marketError.Message = "所选城市不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("所选城市字段长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Region" {
		switch err.Tag() {
		case "required":
			marketError.Message = "所在地区不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("所在地区字段长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Street" {
		switch err.Tag() {
		case "required":
			marketError.Message = "所在街道不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("所在街道字段长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "FullAddress" {
		switch err.Tag() {
		case "required":
			marketError.Message = "详细地址不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("详细地址字段长度不能超过%s个字符", err.Param())
		}
	} else if err.Field() == "Tag" {
		switch err.Tag() {
		case "max":
			marketError.Message = fmt.Sprintf("地址标签字段不能超过%s个字符", err.Param())
		}

	} else if err.Field() == "IsDefault" {
		switch err.Tag() {
		case "required":
			marketError.Message = "是否设置默认地址不能为空"
		}
	}
	return marketError
}

type AddressURISchema struct {
	AddressID uint `uri:"address_id" validate:"required"`
}

func (a AddressURISchema) Validate(errs validator.ValidationErrors) ecode.MarketError {
	var marketError ecode.MarketError
	err := errs[0]

	if err.Field() == "AddressID" {
		switch err.Tag() {
		case "required":
			marketError.Message = "uri地址有误"
		}
	}
	return marketError
}

type AddressSchema struct {
	ID     uint `json:"id"`
	UserID uint `json:"user_id"` // 用户id    用户1 --> 地址N
	AddressBaseSchema
}
