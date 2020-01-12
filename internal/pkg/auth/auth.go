package auth

import (
	"github.com/gin-gonic/gin"

	"market/internal/pkg/auth/jwtauth"
)

type Auth interface {
	// GenerateToken 生成令牌
	GenerateToken(userID uint, username string) (string, error)

	// ParseToken 解析令牌
	ParseToken(token string) (*jwtauth.Claims, error)

	// ParseUserID 从上下文中获取UserID
	ParseUserID(ctx *gin.Context) (uint, error)
}
