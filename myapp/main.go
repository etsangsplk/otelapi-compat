package main

import (
	"context"
	"io/ioutil"
	"net/http"

	v2sdk "github.com/tigrannajaryan/otelapi-compat/opentelemetry/sdk/v2"
	v2 "github.com/tigrannajaryan/otelapi-compat/opentelemetry/v2"
	"github.com/tigrannajaryan/otelapi-compat/third-party/database"
	"github.com/tigrannajaryan/otelapi-compat/third-party/httpserver"
)

func main() {
	v2sdk.EnableSDK()

	srv := &httpserver.Srv{}

	srv.AddHandler("/test", myHandler)
	srv.Start()
}

func myHandler(ctx context.Context) {
	// Here ctx contains a root span created by OpenTelemetry API V1.
	// However we are calling OpenTelemetry API V2 and seamlessly getting
	// getting a child span as opentelemetryv2.Span type.

	_, childSpan := v2.StartSpan(ctx, "client-request")
	defer childSpan.Finish(nil)

	resp, err := http.Get("http://example.com/")
	if err != nil {
		// handle error
	}
	defer resp.Body.Close()
	ioutil.ReadAll(resp.Body)

	database.ExecQuery(ctx, "SELECT * FROM opentelemetry")
}
