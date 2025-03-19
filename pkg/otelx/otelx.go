package otelx

import (
	"context"
	"fmt"
	"portto/pkg/contextx"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/exporters/otlp/otlpmetric/otlpmetricgrpc"
	"go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	"go.opentelemetry.io/otel/propagation"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.10.0"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// SDK is the OpenTelemetry SDK.
type SDK struct {
	target      string
	serviceName string
	verbose     bool
}

// SetupSDK creates a new OpenTelemetry SDK.
func SetupSDK(target string, name string, verbose bool) (*SDK, func(), error) {
	ctx := contextx.WithContext(context.Background())

	instance := &SDK{
		target:      target,
		serviceName: name,
		verbose:     verbose,
	}

	clean, err := instance.setupOTelSDK(ctx)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to setup OpenTelemetry SDK: %w", err)
	}

	return instance, clean, nil
}

func (x *SDK) setupOTelSDK(ctx contextx.Contextx) (func(), error) {
	ctx.Info(
		"setting up OpenTelemetry SDK",
		"service_name", x.serviceName,
		"otlp", x.target,
	)

	var shutdownFuncs []func(context.Context) error

	res, err := resource.New(ctx, resource.WithAttributes(semconv.ServiceNameKey.String(x.serviceName)))
	if err != nil {
		ctx.Error("failed to create resource", "error", err)
		return nil, err
	}

	var conn *grpc.ClientConn
	if x.target != "" {
		conn, err = initConn(x.target)
		if err != nil {
			ctx.Error("failed to create gRPC client", "error", err)
			return nil, err
		}
	}

	tracerProvider, err := x.newTracer(ctx, res, conn)
	if err != nil {
		ctx.Error("failed to create the Jaeger exporter", "error", err)
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, tracerProvider.Shutdown)

	meterProvider, err := x.newMeter(ctx, res, conn)
	if err != nil {
		ctx.Error("failed to create the OTLP exporter", "error", err)
		return nil, err
	}
	shutdownFuncs = append(shutdownFuncs, meterProvider.Shutdown)

	return func() {
		ctx.Info("shutting down OpenTelemetry SDK")
		for _, fn := range shutdownFuncs {
			_ = fn(ctx)
		}
	}, nil
}

func initConn(target string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(target, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to create gRPC client: %w", err)
	}

	return conn, nil
}

func (x *SDK) newTracer(
	c context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
) (*sdktrace.TracerProvider, error) {
	var exporter sdktrace.SpanExporter
	var err error
	if conn == nil && x.verbose {
		exporter, err = stdouttrace.New()
		if err != nil {
			return nil, fmt.Errorf("failed to create stdouttrace: %w", err)
		}
	} else {
		exporter, err = otlptracegrpc.New(c, otlptracegrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("failed to create the Jaeger exporter: %w", err)
		}
	}

	processor := sdktrace.NewBatchSpanProcessor(exporter)
	provider := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithResource(res),
		sdktrace.WithSpanProcessor(processor),
	)
	otel.SetTracerProvider(provider)

	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(
		propagation.TraceContext{},
		propagation.Baggage{},
	))

	return provider, nil
}

func (x *SDK) newMeter(
	ctx context.Context,
	res *resource.Resource,
	conn *grpc.ClientConn,
) (p *sdkmetric.MeterProvider, err error) {
	var exporter sdkmetric.Exporter
	if conn == nil && x.verbose {
		exporter, err = stdoutmetric.New()
		if err != nil {
			return nil, fmt.Errorf("failed to create stdoutmetric: %w", err)
		}
	} else {
		exporter, err = otlpmetricgrpc.New(ctx, otlpmetricgrpc.WithGRPCConn(conn))
		if err != nil {
			return nil, fmt.Errorf("failed to create the OTLP exporter: %w", err)
		}
	}

	provider := sdkmetric.NewMeterProvider(
		sdkmetric.WithReader(sdkmetric.NewPeriodicReader(exporter, sdkmetric.WithInterval(3*time.Second))),
		sdkmetric.WithResource(res),
	)
	otel.SetMeterProvider(provider)

	return provider, nil
}
