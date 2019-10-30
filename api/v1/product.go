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
	)

	if err := c.ShouldBindQuery(&catQuery); err != nil {
		errs := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(http.StatusBadRequest, catQuery.Validate(errs))
		return
	} else {
		params := map[string]interface{}{
			"parent_id": catQuery.ParentId,
		}

		if categories, err = models.GetCategories(params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
			return
		}
		if total, err = models.GetCategoryCount(params); err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
				"message": err.Error(),
			})
		}

		c.JSON(http.StatusOK, gin.H{
			"categories": categories,
			"total":      total,
		})

	}

}
