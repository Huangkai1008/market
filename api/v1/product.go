package v1

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/go-playground/validator.v9"
	"market/pkg/validate"
	"net/http"
)

func GetCategories(c *gin.Context) {
	/**
	获取分类信息
	parent_id 父级id
	*/
	var catQuery validate.CategoryQuery
	if err := c.ShouldBindQuery(&catQuery); err != nil {
		errs := err.(validator.ValidationErrors)
		c.AbortWithStatusJSON(http.StatusBadRequest, catQuery.Validate(errs))
	}

}
