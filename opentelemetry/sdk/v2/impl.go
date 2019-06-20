package v2

import (
	"context"
	"log"
	"time"

	v1api "github.com/tigrannajaryan/otelapi-compat/opentelemetry"
	"github.com/tigrannajaryan/otelapi-compat/opentelemetry/sdk"
	v2api "github.com/tigrannajaryan/otelapi-compat/opentelemetry/v2"
)

// EnableSDK enables SDK V2.
func EnableSDK() {
	impl1 := &ImplV1{}
	impl2 := &ImplV2{}
	v1api.RegisterSDK(impl1)
	v2api.RegisterSDK(impl2)
}

// ImplV1 implements API V1 as a shim over SDK V2.
type ImplV1 struct {
}

// StartSpan starts a new child span of the current span in the context.
// If there is no span in the context, creates a new trace and span.
// Returned context contains the newly created span.
// You can use it to propagate the returned span in process.
func (impl *ImplV1) StartSpan(ctx context.Context, operation string) (context.Context, v1api.Span) {
	childCtx, span := v2api.GetRegisteredSDK().StartSpan(ctx, operation)
	return childCtx, &SpanV1{impl: span}
}

// FromContext returns the Span stored in a context, or nil if there isn't one.
func (impl *ImplV1) FromContext(ctx context.Context) v1api.Span {
	span := v2api.GetRegisteredSDK().FromContext(ctx)
	if span == nil {
		return nil
	}
	return &SpanV1{impl: span}
}

// SpanV1 is implementation of V1 via V2.
type SpanV1 struct {
	impl v2api.Span
}

// Finish the span
func (span *SpanV1) Finish(endTimeUnixEpoch int64) {
	// Here is an example of mapping of incompatible Span types:
	// we convert from integer unix epoch times to time.Time.
	if endTimeUnixEpoch != 0 {
		t := time.Unix(int64(endTimeUnixEpoch), 0)
		span.impl.Finish(&t)
	} else {
		span.impl.Finish(nil)
	}
}

// ImplV2 implements API V2.
type ImplV2 struct {
}

// StartSpan starts a new child span of the current span in the context.
// If there is no span in the context, creates a new trace and span.
// Returned context contains the newly created span.
// You can use it to propagate the returned span in process.
func (impl *ImplV2) StartSpan(ctx context.Context, operation string) (context.Context, v2api.Span) {
	rootSpan := impl.FromContext(ctx)
	span := &SpanImplV2{Parent: rootSpan, Operation: operation, StartTime: time.Now()}
	newCtx := context.WithValue(ctx, sdk.SpanContextKey, span)
	return newCtx, span
}

// FromContext returns the Span stored in a context, or nil if there isn't one.
func (impl *ImplV2) FromContext(ctx context.Context) v2api.Span {
	span := ctx.Value(sdk.SpanContextKey)
	if span == nil {
		return nil
	}
	return span.(v2api.Span)
}

// SpanImplV2 is implementation of V2 span.
type SpanImplV2 struct {
	Parent    v2api.Span
	Operation string
	StartTime time.Time
	EndTime   time.Time
}

// Finish the span
func (span *SpanImplV2) Finish(endTime *time.Time) {
	if endTime != nil {
		span.EndTime = *endTime
	} else {
		span.EndTime = time.Now()
	}
	// Do whatever is needed to finish the span

	log.Printf("Generating span %s, parent %v", span.Operation, span.Parent)
}
