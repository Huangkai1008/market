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

type Register struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
	Email    string `json:"email" validate:"required"`
}

func (r *Register) Validate(errs validator.ValidationErrors) e.MarketError {
	var marketError e.MarketError
	err := errs[0] // 获得第一个错误并返回, yield错误
	fmt.Println(err.Field(), err.ActualTag(), err.Namespace(), err.Param(),
		err.StructField(), err.StructNamespace(), err.Tag(), err.Value())

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
	} else if err.Field() == "Email" {
		switch err.Tag() {
		case "required":
			marketError.Message = "密码不能为空"
		}
	}

	return marketError
}
