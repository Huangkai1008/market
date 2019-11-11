package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/models"
	"market/pkg/validate"
	"net/http"
)

func GetCategories(ctx *gin.Context) {
	/**
	获取分类信息
	parent_id 父级id
	*/
	var catQuery validate.CategoryQuery
	var (
		categories []models.ProductCategory
		total      int
		err        error
	)

	if err := ctx.ShouldBindQuery(&catQuery); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusUnprocessableEntity, catQuery.Validate(errs))
		return
	}

	condition := make(map[string]interface{})
	condition["parent_id"] = catQuery.ParentId

	if categories, err = models.GetCategories(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if total, err = models.GetCategoryCount(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"total":      total,
	})
}
