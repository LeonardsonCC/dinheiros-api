package logger

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hako/durafmt"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// add logger to request
		logger := newRequestLogger()
		logger = logger.With().Str("request-id", uuid.NewString()).Logger()
		ctx := logger.WithContext(c.Request.Context())

		// updates request context with the logger
		c.Request = c.Request.Clone(ctx)

		start := time.Now()
		logger.Info().Msg("request started")
		c.Next()

		duration := time.Since(start)
		logger.
			Info().
			Str("duration-pretty", durafmt.Parse(duration).String()).
			Int64("duration-ms", duration.Milliseconds()).
			Msg("request finished")
	}
}
