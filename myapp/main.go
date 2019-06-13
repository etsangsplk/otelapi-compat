package main

import (
	"context"
	"io/ioutil"
	"net/http"

	"github.com/tigrannajaryan/otelapi-compat/httpserver"
	"github.com/tigrannajaryan/otelapi-compat/opentelemetryv2"
)

func main() {
	srv := &httpserver.Srv{}

	srv.AddHandler("/test", myHandler)
	srv.Start()
}

func myHandler(ctx context.Context) {
	// Here ctx contains a root span created by OpenTelemetry API V1.
	// However we are calling OpenTelemetry API V2 and seamlessly getting
	// getting a child span as opentelemetryv2.Span type.

	_, childSpan := opentelemetryv2.StartSpan(ctx, "client-request")
	defer childSpan.Finish(nil)

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)
}
