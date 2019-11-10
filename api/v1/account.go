package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/models"
	"market/pkg/validate"
	"net/http"
)

func CreateAddress(ctx *gin.Context) {
	/**
	创建收货地址
	*/

	var addressSchema validate.AddressSchema

	if err := ctx.ShouldBindJSON(&addressSchema); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, addressSchema.Validate(errs))
		return
	}

	userId := uint(ctx.GetInt("userId"))

	address := models.Address{
		UserID:      userId,
		Consignee:   addressSchema.Consignee,
		Mobile:      addressSchema.Mobile,
		Region:      addressSchema.Region,
		FullAddress: addressSchema.FullAddress,
		Tag:         addressSchema.Tag,
		IsDefault:   &addressSchema.IsDefault,
	}

	if address, err := models.CreateAddressTx(address); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, address)
	}
}

func GetAddresses(ctx *gin.Context) {
	/**
	获取收货地址
	*/

	maps := make(map[string]interface{})
	if addresses, err := models.GetAddresses(maps); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, addresses)
	}

}
