package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"time"
)

func GinZap(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)
		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				logger.Warn("[GIN]",
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					zap.String("method", method),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("error", c.Errors.String()),
				)
			}
		case statusCode >= 500:
			{
				logger.Error("[GIN]",
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					zap.String("method", method),
					zap.String("user-agent", c.Request.UserAgent()),
					zap.String("error", c.Errors.String()),
				)
			}
		default:
			logger.Info("[GIN]",
				zap.Int("statusCode", statusCode),
				zap.String("latency", latency.String()),
				zap.String("clientIP", clientIP),
				zap.String("method", method),
				zap.String("user-agent", c.Request.UserAgent()),
				zap.String("error", c.Errors.String()),
			)
		}
	}

}
