package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/pkg/validate"
	"net/http"
)

func Register(c *gin.Context) {
	/**
	用户注册
	username 用户名
	password 密码
	email 邮箱
	*/
	var register validate.Register
	if err := c.ShouldBindJSON(&register); err != nil {
		errs := err.(validator.ValidationErrors)
		c.JSON(http.StatusBadRequest, register.Validate(errs))
	}
}
