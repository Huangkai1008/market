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

// Register 用户注册
func (h *Handler) Register(ctx *gin.Context) {
	var registerSchema RegisterSchema

	if err := ctx.ShouldBindJSON(&registerSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, registerSchema.Validate(errs))
		return
	}

	condition := make(map[string]interface{})

	condition["username"] = registerSchema.Username
	if exist, err := h.repository.ExistUser(condition); exist || (err != nil) {
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
	if exist, err := h.repository.ExistUser(condition); exist || (err != nil) {
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

	if user, err := h.repository.CreateUser(&user); err != nil {
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
	user, err := h.repository.GetUser(condition)
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

// GetAddress 获取收货地址
func (h *Handler) GetAddress(ctx *gin.Context) {
	userId, err := h.auth.ParseUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	condition := make(map[string]interface{})
	condition["user_id"] = userId

	if addresses, err := h.repository.GetAddresses(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, addresses.ToSchemaAddresses())
	}
}

// CreateAddress 创建收货地址
func (h *Handler) CreateAddress(ctx *gin.Context) {
	var addressCreate AddressCreateSchema

	userId, err := h.auth.ParseUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindJSON(&addressCreate); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, addressCreate.Validate(errs))
		return
	}

	address := addressCreate.ToAddress()
	address.UserID = userId

	if address, err := h.repository.CreateAddress(address); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, address.ToSchemaAddress())
	}
}

// UpdateAddress 更新收货地址
func (h *Handler) UpdateAddress(ctx *gin.Context) {
	var (
		addressURI    AddressURISchema
		addressUpdate AddressUpdateSchema
	)

	userId, err := h.auth.ParseUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindUri(&addressURI); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, addressURI.Validate(errs))
		return
	}

	if err := ctx.ShouldBindJSON(&addressUpdate); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, addressUpdate.Validate(errs))
		return
	}

	addressID := addressURI.AddressID

	maps := addressUpdate.ToMap()

	if address, err := h.repository.UpdateAddress(addressID, userId, maps); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, (&address).ToSchemaAddress())
	}

}

// DeleteAddress 删除收货地址
func (h *Handler) DeleteAddress(ctx *gin.Context) {
	var addressURI AddressURISchema

	userId, err := h.auth.ParseUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindUri(&addressURI); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, addressURI.Validate(errs))
		return
	}
	addressID := addressURI.AddressID

	condition := make(map[string]interface{})
	condition["user_id"] = userId
	condition["id"] = addressID
	if err := h.repository.DeleteAddress(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
