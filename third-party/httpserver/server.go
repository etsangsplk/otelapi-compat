package httpserver

import (
	"context"
	"net/http"

	"github.com/tigrannajaryan/otelapi-compat/opentelemetry"
)

// Srv is my super server
type Srv struct {
}

// AddHandler to the server.
func (srv *Srv) AddHandler(route string, handler func(ctx context.Context)) {
	http.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
		// Here we are using OpenTelementry API V1 and creating a span V1 in the context.
		ctx, span := opentelemetry.StartSpan(context.Background(), route)
		defer span.Finish(0)

		// pass the context to the handler.
		handler(ctx)
	})
}

// Start the server.
func (srv *Srv) Start() {
	http.ListenAndServe("0.0.0.0:9988", nil)
}
