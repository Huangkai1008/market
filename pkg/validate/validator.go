package validate

import (
	"fmt"
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
	"market/pkg/e"
)

func init() {
	binding.Validator = new(defaultValidator)
}

/************************************/
/**********   用户模块验证    ********/
/************************************/

type Register struct {
	Username string `json:"username" validate:"required,max=100"`
	Password string `json:"password" validate:"required,max=100"`
	Email    string `json:"email" validate:"required,email,max=256"`
}

type Login struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}

func (r *Register) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
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

func (l *Login) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
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

/************************************/
/**********   账户模块验证    ********/
/************************************/
type AddressSchema struct {
	Consignee   string `json:"consignee" validate:"required,max=64"`
	Mobile      string `json:"mobile" validate:"required,max=32"`
	Region      string `json:"region" validate:"required,max=32"`
	FullAddress string `json:"full_address" validate:"required,max=64"`
	Tag         string `json:"tag" validate:"max=32"`
	IsDefault   *bool  `json:"is_default" validate:"required"`
}
type (
	//AddressCreate 收货地址创建schema
	AddressCreate struct {
		AddressSchema
	}

	//AddressUpdate 收货地址更新schema
	AddressUpdate struct {
		AddressSchema
	}
)

func (a *AddressSchema) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
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
	} else if err.Field() == "Region" {
		switch err.Tag() {
		case "required":
			marketError.Message = "所在地区不能为空"
		case "max":
			marketError.Message = fmt.Sprintf("所在地区字段长度不能超过%s个字符", err.Param())
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

type AddressUri struct {
	AddressID uint `uri:"address_id" validate:"required"`
}

func (a AddressUri) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
	err := errs[0]

	if err.Field() == "AddressID" {
		switch err.Tag() {
		case "required":
			marketError.Message = "uri地址有误"
		}
	}
	return marketError
}

/************************************/
/**********   商品模块验证    ********/
/************************************/
type CategoryQuery struct {
	ParentId string `form:"parent_id" validate:"required,gte=0"`
}

func (c *CategoryQuery) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
	err := errs[0]

	if err.Field() == "ParentId" {
		switch err.Tag() {
		case "required":
			marketError.Message = "分类父级id不能为空"
		case "gte":
			marketError.Message = "分类父级id不能小于0"
		}
	}
	return marketError
}
