package logger

import (
	"context"
	"os"

	"github.com/rs/zerolog"
)

func FromContext(ctx context.Context) *zerolog.Logger {
	return zerolog.Ctx(ctx)
}

func newRequestLogger() zerolog.Logger {
	return zerolog.New(os.Stdout).
		With().
		Timestamp().
		Logger()
}
