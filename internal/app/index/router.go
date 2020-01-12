package index

import (
	"github.com/gin-gonic/gin"

	"market/internal/pkg/router"
)

func NewRouter(
	h *Handler,
) router.Router {
	return func(r *gin.Engine) {
		indexApi := r.Group("/")
		{
			indexApi.GET("ping", Ping)
		}
	}
}
