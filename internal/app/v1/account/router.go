package account

import (
	"github.com/gin-gonic/gin"

	"market/internal/pkg/middleware"
	"market/internal/pkg/router"
)

func NewRouter(
	h *Handler,
) router.Group {
	return func(r *gin.RouterGroup) {
		accountApi := r.Group("/account", middleware.AuthMiddleware(h.auth))
		{
			accountApi.GET("addresses", h.Get)
			accountApi.POST("addresses", h.Create)
			accountApi.PUT("addresses/:address_id", h.Update)
			accountApi.DELETE("addresses/:address_id", h.Delete)
		}
	}
}
