package database

import (
	"context"

	"github.com/tigrannajaryan/otelapi-compat/opentelemetryv1"
)

// ExecQuery executes a database query.
func ExecQuery(ctx context.Context, sql string) {
	_, span := opentelemetryv1.StartSpan(ctx, "database-call")
	defer span.Finish(0)

	// Do the query execution
}
