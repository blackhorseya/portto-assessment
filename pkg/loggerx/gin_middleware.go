package loggerx

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
	"portto/pkg/contextx"
)

// GinTraceLoggingMiddleware attaches the provided logger to the request context
// and logs the request start and completion details including duration.
func GinTraceLoggingMiddleware(logger *slog.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Attach the logger to the request context.
		newCtx := contextx.WithLogger(c.Request.Context(), logger)
		c.Request = c.Request.WithContext(newCtx)

		// Log request start.
		logger.Info("request started", "path", c.Request.URL.Path, "method", c.Request.Method)

		start := time.Now()
		c.Next()
		duration := time.Since(start)

		// Log request completion.
		logger.Info("request completed",
			"status", c.Writer.Status(),
			"path", c.Request.URL.Path,
			"duration", duration.String(),
		)
	}
}
