package sdk

import (
	"context"
	"time"

	"github.com/tigrannajaryan/otelapi-compat/opentelemetry"
)

// SpanImpl V1
type SpanImpl struct {
	Parent    opentelemetry.Span
	Operation string
	StartTime int64
	EndTime   int64
}

// Impl V1
type Impl struct {
}

type opentelemetryContextKey string

// SpanContextKey is the key of span value in the context.
const SpanContextKey opentelemetryContextKey = "span"

// StartSpan starts a new child span of the current span in the context.
// If there is no span in the context, creates a new trace and span.
// Returned context contains the newly created span.
// You can use it to propagate the returned span in process.
func (impl *Impl) StartSpan(ctx context.Context, operation string) (context.Context, opentelemetry.Span) {
	rootSpan := FromContext(ctx)
	span := &SpanImpl{Parent: rootSpan, Operation: operation, StartTime: time.Now().Unix()}
	newCtx := context.WithValue(ctx, SpanContextKey, span)
	return newCtx, span
}

// Finish the span
func (span *SpanImpl) Finish(endTimeUnixEpoch int64) {
	span.EndTime = endTimeUnixEpoch
}

// FromContext returns the Span stored in a context, or nil if there isn't one.
func FromContext(ctx context.Context) opentelemetry.Span {
	span := ctx.Value(SpanContextKey)
	if span == nil {
		return nil
	}
	return span.(opentelemetry.Span)
}
