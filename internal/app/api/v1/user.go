package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"market/internal/app/model"
	"market/internal/app/schema"
	"market/internal/app/validate"
	"market/pkg/utils"
)

func Register(ctx *gin.Context) {
	/**
	用户注册
	username 用户名
	password 密码
	email 邮箱
	*/
	var register validate.Register

	if err := ctx.ShouldBindJSON(&register); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, register.Validate(errs))
		return
	}

	condition := make(map[string]interface{})

	condition["username"] = register.Username
	if exist, err := model.ExistUser(condition); exist || (err != nil) {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if exist {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "存在相同的用户名",
			})
			return
		}
	}

	delete(condition, "username")
	condition["email"] = register.Email
	if exist, err := model.ExistUser(condition); exist || (err != nil) {
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if exist {
			ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": "存在相同的邮箱账户",
			})
			return
		}
	}

	user := model.User{
		Username:     register.Username,
		Email:        register.Email,
		HashPassword: utils.MD5(register.Password),
	}

	if user, err := model.CreateUser(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, user.ToSchemaUser())
	}

}

func GetToken(ctx *gin.Context) {
	/**
	已注册用户获取token
	*/
	var login validate.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, login.Validate(errs))
		return
	}

	condition := make(map[string]interface{})

	condition["username"] = login.Username
	user, err := model.GetUser(condition)
	if gorm.IsRecordNotFoundError(err) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "不存在的用户名",
		})
		return
	}
	if utils.MD5(login.Password) != user.HashPassword {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "用户名和密码不匹配",
		})
		return
	}

	if token, err := utils.GenerateToken(user.ID, user.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "token生成错误",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, &schema.TokenBack{Token: token})
	}
}
