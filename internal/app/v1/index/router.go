package index

import (
	"github.com/gin-gonic/gin"
)

func NewRouter(
	h *Handler,
) func(r *gin.RouterGroup) {
	return func(r *gin.RouterGroup) {
		indexApi := r.Group("/")
		{
			indexApi.GET("ping", Ping)
		}
	}
}
