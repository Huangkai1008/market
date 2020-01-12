package middleware

import (
	"bytes"
	"time"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"

	"market/internal/pkg/ecode"
)

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// LoggerMiddleware 日志中间件
func LoggerMiddleware(logger *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		bodyLogWriter := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = bodyLogWriter
		c.Next()
		latency := time.Since(start)

		responseBody := bodyLogWriter.body.String()

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()
		requestUri := c.Request.RequestURI

		switch {
		case statusCode >= 400 && statusCode <= 499:
			{
				logger.Warn("[WARN]",
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					zap.String("method", method),
					zap.String("requestUri", requestUri),
					zap.String("error", ecode.GetMsg(responseBody)),
				)
			}
		case statusCode >= 500:
			{
				logger.Error("[ERROR]",
					zap.Int("statusCode", statusCode),
					zap.String("latency", latency.String()),
					zap.String("clientIP", clientIP),
					zap.String("method", method),
					zap.String("requestUri", requestUri),
					zap.String("error", ecode.GetMsg(responseBody)),
				)
			}
		default:
			logger.Info("[INFO]",
				zap.Int("statusCode", statusCode),
				zap.String("latency", latency.String()),
				zap.String("clientIP", clientIP),
				zap.String("method", method),
				zap.String("requestUri", requestUri),
			)
		}
	}
}
