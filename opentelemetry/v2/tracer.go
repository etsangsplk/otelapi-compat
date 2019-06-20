package v2

import (
	"context"
	"time"
)

// Span V2
type Span interface {
	Finish(endTimeUnixEpoch *time.Time)
}

// SDK V2
type SDK interface {
	StartSpan(ctx context.Context, operation string) (context.Context, Span)
	FromContext(ctx context.Context) Span
}

var registeredSDK SDK

// RegisterSDK V2
func RegisterSDK(sdk SDK) {
	registeredSDK = sdk
}

// GetRegisteredSDK V2
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
