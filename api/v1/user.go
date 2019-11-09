package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/models"
	"market/pkg/utils"
	"market/pkg/validate"
	"net/http"
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, register.Validate(errs))
		return
	}

	maps := make(map[string]interface{})

	maps["username"] = register.Username
	if exist, err := models.ExistUser(maps); exist || (err != nil) {
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

	delete(maps, "username")
	maps["email"] = register.Email
	if exist, err := models.ExistUser(maps); exist || (err != nil) {
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

	user := models.User{
		Username:     register.Username,
		Email:        register.Email,
		HashPassword: utils.MD5(register.Password),
	}

	if user, err := models.CreateUser(user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, user)
	}

}

func GetToken(ctx *gin.Context) {
	/**
	已注册用户获取token
	*/
	var login validate.Login

	if err := ctx.ShouldBindJSON(&login); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, login.Validate(errs))
		return
	}

	maps := make(map[string]interface{})

	maps["username"] = login.Username
	user, err := models.GetUser(maps)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if user == (models.User{}) {
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
		ctx.JSON(http.StatusOK, gin.H{
			"token": token,
		})
	}

}
