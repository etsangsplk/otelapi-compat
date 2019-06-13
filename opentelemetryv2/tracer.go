package opentelemetryv2

import (
	"context"
	"log"
	"time"
)

// Span v2
type Span struct {
	Parent    *Span
	Operation string
	StartTime time.Time
	EndTime   time.Time
}

type opentelemetryContextKey string

const spanContextKey opentelemetryContextKey = "span"

// StartSpan starts a new child span of the current span in the context.
// If there is no span in the context, creates a new trace and span.
// Returned context contains the newly created span.
// You can use it to propagate the returned span in process.
func StartSpan(ctx context.Context, operation string) (context.Context, *Span) {
	rootSpan := FromContext(ctx)
	span := &Span{Parent: rootSpan, Operation: operation, StartTime: time.Now()}
	newCtx := context.WithValue(ctx, spanContextKey, span)
	return newCtx, span
}

// Finish the span
func (span *Span) Finish(endTime *time.Time) {
	if endTime != nil {
		span.EndTime = *endTime
	} else {
		span.EndTime = time.Now()
	}
	// Do whatever is needed to finish the span

	log.Printf("Generating span %s, parent %v", span.Operation, span.Parent)
}

// FromContext returns the Span stored in a context, or nil if there isn't one.
func FromContext(ctx context.Context) *Span {
	span := ctx.Value(spanContextKey)
	if span == nil {
		return nil
	}
	return span.(*Span)
}
