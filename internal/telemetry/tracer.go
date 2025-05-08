package telemetry

import (
	"context"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

var tracer trace.Tracer

func GetAppTracer() trace.Tracer {
	if tracer == nil {
		tracer = otel.Tracer("dinheiros-api")
	}
	return tracer
}

func SpanFromContext(ctx context.Context) trace.Span {
	return trace.SpanFromContext(ctx)
}
