package logger

import (
	"time"

	"github.com/LeonardsonCC/dinheiros/internal/telemetry"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/hako/durafmt"
)

func Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// add logger to request
		logger := newRequestLogger()
		logger = logger.With().Str("request-id", uuid.NewString()).Logger()

		telemetrySpan := telemetry.SpanFromContext(c.Request.Context())
		logger = logger.With().Str(
			"traceID",
			telemetrySpan.SpanContext().TraceID().String(),
		).Logger()

		ctx := logger.WithContext(c.Request.Context())

		// updates request context with the logger
		c.Request = c.Request.Clone(ctx)

		start := time.Now()
		logger.
			Info().
			Str("request.path", c.Request.URL.Path).
			Msg("request started")
		c.Next()

		duration := time.Since(start)
		logger.
			Info().
			Str("response.duration_pretty", durafmt.Parse(duration).String()).
			Int64("response.duration_ms", duration.Milliseconds()).
			Int("response.status_code", c.Writer.Status()).
			Msg("request finished")
	}
}
