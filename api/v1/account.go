package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/models"
	"market/pkg/utils"
	"market/pkg/validate"
	"net/http"
)

func GetAddresses(ctx *gin.Context) {
	/**
	获取收货地址
	*/

	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	condition := make(map[string]interface{})
	condition["user_id"] = userId

	if addresses, err := models.GetAddresses(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, addresses.ToSchemaAddresses())
	}

}

func CreateAddress(ctx *gin.Context) {
	/**
	创建收货地址
	*/

	var addressCreate validate.AddressCreate

	userId, err := utils.GetUserID(ctx)
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

	address := models.Address{
		UserID:      userId,
		Consignee:   addressCreate.Consignee,
		Mobile:      addressCreate.Mobile,
		Region:      addressCreate.Region,
		FullAddress: addressCreate.FullAddress,
		Tag:         addressCreate.Tag,
		IsDefault:   addressCreate.IsDefault,
	}

	if address, err := models.CreateAddressTx(&address); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusCreated, address.ToSchemaAddress())
	}
}

func UpdateAddress(ctx *gin.Context) {
	/**
	修改收货地址
	*/

	var (
		addressUri    validate.AddressUri
		addressUpdate validate.AddressUpdate
	)

	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindUri(&addressUri); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, addressUri.Validate(errs))
		return
	}

	if err := ctx.ShouldBindJSON(&addressUpdate); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, addressUpdate.Validate(errs))
		return
	}

	addressID := addressUri.AddressID

	maps := make(map[string]interface{})
	maps["consignee"] = addressUpdate.Consignee
	maps["mobile"] = addressUpdate.Mobile
	maps["region"] = addressUpdate.Region
	maps["full_address"] = addressUpdate.Consignee
	maps["tag"] = addressUpdate.Tag
	maps["is_default"] = *addressUpdate.IsDefault

	if address, err := models.UpdateAddressTx(addressID, userId, maps); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, (&address).ToSchemaAddress())
	}

}

func DeleteAddress(ctx *gin.Context) {
	/**
	删除收货地址
	*/
	var addressUri validate.AddressUri

	userId, err := utils.GetUserID(ctx)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	if err := ctx.ShouldBindUri(&addressUri); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, addressUri.Validate(errs))
		return
	}
	addressID := addressUri.AddressID

	condition := make(map[string]interface{})
	condition["user_id"] = userId
	condition["id"] = addressID
	if err := models.DeleteAddress(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{})
	}

}
