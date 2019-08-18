package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/models"
	"market/pkg/utils"
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
		c.AbortWithStatusJSON(http.StatusBadRequest, register.Validate(errs))
	} else {
		params := map[string]interface{}{
			"username": register.Username,
		}
		if exist := models.ExistUser(params); exist {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "存在相同的用户名",
			})
			return
		}
		params = map[string]interface{}{
			"email": register.Email,
		}
		if exist := models.ExistUser(params); exist {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "存在相同的邮箱账户",
			})
			return
		}

		user := models.User{
			Username:     register.Username,
			Email:        register.Email,
			HashPassword: utils.MD5(register.Password),
		}

		if err, user := models.CreateUser(user); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		} else {
			c.JSON(http.StatusOK, user)
		}
	}
}
