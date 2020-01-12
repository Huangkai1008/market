package user

import (
	"github.com/gin-gonic/gin"

	"market/internal/pkg/router"
)

func NewRouter(
	h *Handler,
) router.Router {
	return func(r *gin.Engine) {
		userApi := r.Group("/users")
		{
			userApi.POST("register", h.Register)
			userApi.POST("tokens", h.Login)
		}
	}
}
