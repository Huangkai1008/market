package index

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"market/internal/pkg/auth"
)

type Handler struct {
	repository Repository
	auth       auth.Auth
}

func NewHandler(repo Repository, auth auth.Auth) *Handler {
	return &Handler{repository: repo, auth: auth}
}

func Ping(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{"msg": "ping"})
}
