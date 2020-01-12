package product

import (
	"github.com/gin-gonic/gin"

	"market/internal/pkg/middleware"
	"market/internal/pkg/router"
)

func NewRouter(
	h *Handler,
) router.Router {
	return func(r *gin.Engine) {
		categoryApi := r.Group("/categories", middleware.AuthMiddleware(h.auth))
		{
			categoryApi.GET("", h.GetCategories)
			categoryApi.GET("/:cat_id/specs", h.GetCategorySpecs)
		}
	}
}
