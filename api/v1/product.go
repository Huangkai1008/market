package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/models"
	"market/pkg/validate"
	"net/http"
)

func GetCategories(c *gin.Context) {
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

	if err := c.ShouldBindQuery(&catQuery); err != nil {
		errs := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(http.StatusBadRequest, catQuery.Validate(errs))
		return
	}

	maps := make(map[string]interface{})
	maps["parent_id"] = catQuery.ParentId

	if categories, err = models.GetCategories(maps); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if total, err = models.GetCategoryCount(maps); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"categories": categories,
		"total":      total,
	})
}
