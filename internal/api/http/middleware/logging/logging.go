package logging

import (
	"time"

	"github.com/antongoncharik/crypto-knight-api/internal/logger"
	"github.com/gin-gonic/gin"
)

func UseLogging() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()

		c.Next()

		latency := time.Since(start)

		logger.Log.Infow(
			"http_request",
			"method", c.Request.Method,
			"path", c.Request.URL.Path,
			"status", c.Writer.Status(),
			"latency", latency.String(),
		)
	}
}
