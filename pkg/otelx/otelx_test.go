package otelx

import (
	"context"
	"testing"
	"time"

	"go.opentelemetry.io/otel"
)

func TestSetupSDK(t *testing.T) {
	target := ""
	name := "test-service"

	sdk, cleanup, err := SetupSDK(target, name, true)
	if err != nil {
		t.Fatalf("failed to setup SDK: %v", err)
	}
	defer cleanup()

	if sdk == nil {
		t.Fatal("expected SDK to be non-nil")
	}

	_, span := otel.GetTracerProvider().Tracer(name).Start(context.Background(), "test")
	defer span.End()

	time.Sleep(1 * time.Second)
}
