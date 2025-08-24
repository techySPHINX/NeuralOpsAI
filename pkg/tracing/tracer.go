package tracing

import (
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// NewTracerProvider creates a new no-op tracer provider.
func NewTracerProvider() trace.TracerProvider {
	return trace.NewNoopTracerProvider()
}

// InitTracer initializes the global tracer provider.
func InitTracer() {
	otel.SetTracerProvider(NewTracerProvider())
}
