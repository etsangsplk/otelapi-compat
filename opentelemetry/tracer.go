// Package opentelemetry is V1 of API.
package opentelemetry

import (
	"context"
)

// Span V1
type Span interface {
	Finish(endTimeUnixEpoch int64)
}

// SDK V1
type SDK interface {
	StartSpan(ctx context.Context, operation string) (context.Context, Span)
	FromContext(ctx context.Context) Span
}

var registeredSDK SDK

// RegisterSDK V1
func RegisterSDK(sdk SDK) {
	registeredSDK = sdk
}

// GetRegisteredSDK V1
func GetRegisteredSDK() SDK {
	return registeredSDK
}

// StartSpan starts span.
func StartSpan(ctx context.Context, operation string) (context.Context, Span) {
	return GetRegisteredSDK().StartSpan(ctx, operation)
}

// FromContext retusn span.
func FromContext(ctx context.Context) Span {
	return GetRegisteredSDK().FromContext(ctx)
}
