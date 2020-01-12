package account

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"market/internal/pkg/auth"
)

type Handler struct {
	repository Repository
	auth       auth.Auth
}

func NewHandler(repo Repository, auth auth.Auth) *Handler {
	return &Handler{repository: repo, auth: auth}
}

// GetOne 获取收货地址
func (h *Handler) Get(ctx *gin.Context) {
	userId, err := h.auth.ParseUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	condition := make(map[string]interface{})
	condition["user_id"] = userId

	if addresses, err := h.repository.Get(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, addresses.ToSchemaAddresses())
	}
}

// Create 创建收货地址
func (h *Handler) Create(ctx *gin.Context) {
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

	address := Address{
		UserID:      userId,
		Consignee:   addressCreate.Consignee,
		Mobile:      addressCreate.Mobile,
		Region:      addressCreate.Region,
		FullAddress: addressCreate.FullAddress,
		Tag:         addressCreate.Tag,
		IsDefault:   addressCreate.IsDefault,
	}

	if address, err := h.repository.Create(&address); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, address.ToSchemaAddress())
	}
}

// Update 更新收货地址
func (h *Handler) Update(ctx *gin.Context) {
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
		ctx.AbortWithStatusJSON(http.StatusBadRequest, addressUpdate.Validate(errs))
		return
	}

	addressID := addressURI.AddressID

	maps := make(map[string]interface{})
	maps["consignee"] = addressUpdate.Consignee
	maps["mobile"] = addressUpdate.Mobile
	maps["region"] = addressUpdate.Region
	maps["full_address"] = addressUpdate.Consignee
	maps["tag"] = addressUpdate.Tag
	maps["is_default"] = *addressUpdate.IsDefault

	if address, err := h.repository.Update(addressID, userId, maps); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, (&address).ToSchemaAddress())
	}

}

// Delete 删除收货地址
func (h *Handler) Delete(ctx *gin.Context) {
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
	if err := h.repository.Delete(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}
}
