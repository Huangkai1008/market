package user

import (
	"github.com/gin-gonic/gin"
	"market/internal/pkg/middleware"

	"market/internal/pkg/router"
)

func NewRouter(
	h *Handler,
) router.Group {
	return func(r *gin.RouterGroup) {
		userApi := r.Group("/users")
		{
			userApi.POST("register", h.Register)
			userApi.POST("login", h.Login)
			addressApi := userApi.Group("addresses", middleware.AuthMiddleware(h.auth))
			{
				addressApi.GET("", h.GetAddress)
				addressApi.POST("", h.CreateAddress)
				addressApi.PUT(":address_id", h.UpdateAddress)
				addressApi.DELETE(":address_id", h.DeleteAddress)
			}
		}
	}
}
