package handler

import "github.com/gin-gonic/gin"

type RestHandler interface {
	GetAddress(ctx *gin.Context)
	Create(ctx *gin.Context)
	Update(ctx *gin.Context)
	Delete(ctx *gin.Context)
}
