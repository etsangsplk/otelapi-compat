// Package opentelemetry is a shim that provides OpenTelemetry API v1
// by mapping the calls to API v2.
package opentelemetry

import (
	"context"
	"time"

	v2 "github.com/tigrannajaryan/otelapi-compat/opentelemetry/v2"
)

// Span v1
type Span struct {
	impl *v2.Span
}

// StartSpan creates a new child span if context has a root
// span or a child if there is no root.
func StartSpan(ctx context.Context, operation string) (context.Context, *Span) {
	ctx, span := v2.StartSpan(ctx, operation)
	return ctx, &Span{impl: span}
}

// Finish the span
func (span *Span) Finish(endTimeUnixEpoch int) {
	// Here is an example of mapping of incompatible Span types:
	// we convert from integer unix epoch times to time.Time.
	if endTimeUnixEpoch != 0 {
		t := time.Unix(int64(endTimeUnixEpoch), 0)
		span.impl.Finish(&t)
	} else {
		span.impl.Finish(nil)
	}
}

// FromContext returns the Span stored in a context, or nil if there isn't one.
func FromContext(ctx context.Context) *Span {
	span := v2.FromContext(ctx)
	if span == nil {
		return nil
	}
	return &Span{impl: span}
}
