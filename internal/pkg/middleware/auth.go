package middleware

import (
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"

	"market/internal/pkg/auth"
)

// AuthMiddleware 授权中间件
func AuthMiddleware(auth auth.Auth) gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.Request.Header // 获取请求头
		authorization := header.Get("Authorization")
		if authorization == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "未携带token信息, token验证失败",
			})
		} else {
			arr := strings.Split(authorization, " ")
			token := arr[1]

			if claims, err := auth.ParseToken(token); err != nil {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "token验证失败",
				})
			} else if claims.ExpiresAt < time.Now().Unix() {
				c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
					"message": "token已过期, token验证失败",
				})
			} else {
				c.Set("userId", claims.UserId)
				c.Next()
			}
		}
	}
}
