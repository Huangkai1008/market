package product

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

// GetCategories 获取分类信息
func (h *Handler) GetCategories(ctx *gin.Context) {
	var (
		catQuerySchema CategoryQuerySchema
		categories     Categories
		total          int
		err            error
	)

	if err := ctx.ShouldBindQuery(&catQuerySchema); err != nil {
		errs := err.(validator.ValidationErrors)
		ctx.AbortWithStatusJSON(http.StatusBadRequest, catQuerySchema.Validate(errs))
		return
	}

	maps := catQuerySchema.ToMap()
	if categories, err = h.repository.GetCategories(maps); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}
	if total, err = h.repository.GetCategoryCount(maps); err != nil {
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

// GetCategorySpecs 获取分类规格信息
func (h *Handler) GetCategorySpecs(ctx *gin.Context) {

	var categorySpecURISchema CategorySpecURISchema

	if err := ctx.ShouldBindUri(&categorySpecURISchema); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
		})
		return
	}

	condition := make(map[string]interface{})
	condition["cat_id"] = categorySpecURISchema.CatId

	if specs, err := h.repository.GetCategorySpec(condition); err != nil {
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
