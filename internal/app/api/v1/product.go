package v1

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"

	"market/internal/app/model"
	"market/internal/app/validate"
)

func GetCategories(ctx *gin.Context) {
	/**
	获取分类信息
	parent_id 父级id
	*/
	var catQuery validate.CategoryQuery
	var (
		categories model.ProductCategories
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

	if categories, err = model.GetCategories(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if total, err = model.GetCategoryCount(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"categories": categories.ToSchemaCategories(),
		"total":      total,
	})
}

func GetCategorySpecs(ctx *gin.Context) {
	/**
	获取分类规格信息
	*/

	var categorySpecUri validate.CategorySpecUri

	if err := ctx.ShouldBindUri(&categorySpecUri); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	condition := make(map[string]interface{})
	condition["cat_id"] = categorySpecUri.CatId

	if specs, err := model.GetCategorySpecs(condition); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	} else {
		ctx.JSON(http.StatusOK, gin.H{
			"specs": specs.ToSchemaCategorySpecs(),
		})
	}

}
