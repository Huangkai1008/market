package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"gopkg.in/go-playground/validator.v9"

	"market/internal/pkg/auth"
	"market/internal/pkg/utils"
)

type Handler struct {
	repository Repository
	auth       auth.Auth
}

func NewHandler(repo Repository, auth auth.Auth) *Handler {
	return &Handler{repository: repo, auth: auth}
}

// RegisterSchema 用户注册
func (h *Handler) Register(ctx *gin.Context) {
	var registerSchema RegisterSchema

	if err := ctx.ShouldBindJSON(&registerSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, registerSchema.Validate(errs))
		return
	}

	condition := make(map[string]interface{})

	condition["username"] = registerSchema.Username
	if exist, err := h.repository.Exist(condition); exist || (err != nil) {
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
	condition["email"] = registerSchema.Email
	if exist, err := h.repository.Exist(condition); exist || (err != nil) {
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

	user := User{
		Username:     registerSchema.Username,
		Email:        registerSchema.Email,
		HashPassword: utils.MD5(registerSchema.Password),
	}

	if user, err := h.repository.Create(&user); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, user.ToSchemaUser())
	}

}

// Login 已注册用户登录获取token
func (h *Handler) Login(ctx *gin.Context) {
	var loginSchema LoginSchema

	if err := ctx.ShouldBindJSON(&loginSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, loginSchema.Validate(errs))
		return
	}

	condition := make(map[string]interface{})

	condition["username"] = loginSchema.Username
	user, err := h.repository.GetOne(condition)
	if gorm.IsRecordNotFoundError(err) {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "不存在的用户名",
		})
		return
	}
	if utils.MD5(loginSchema.Password) != user.HashPassword {
		ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
			"message": "用户名和密码不匹配",
		})
		return
	}

	if token, err := h.auth.GenerateToken(user.ID, user.Username); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": "token生成错误",
		})
		return
	} else {
		ctx.JSON(http.StatusOK, &TokenBackSchema{Token: token})
	}
}
