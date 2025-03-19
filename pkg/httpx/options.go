package httpx

import (
	"log/slog"
	"time"
)

// ServerOption defines a function type for configuring server options.
type ServerOption func(opts *serverOptions)

// serverOptions holds configuration options for starting and shutting down the server.
type serverOptions struct {
	Host            string
	Port            int
	Logger          *slog.Logger
	ShutdownTimeout time.Duration
}

// WithHost sets the host of the server
func WithHost(host string) ServerOption {
	return func(opts *serverOptions) {
		opts.Host = host
	}
}

// WithPort sets the port of the server
func WithPort(port int) ServerOption {
	return func(opts *serverOptions) {
		opts.Port = port
	}
}

// WithLogger sets the logger of the server
func WithLogger(logger *slog.Logger) ServerOption {
	return func(opts *serverOptions) {
		opts.Logger = logger
	}
}

// WithShutdownTimeout sets the shutdown timeout of the server
func WithShutdownTimeout(timeout time.Duration) ServerOption {
	return func(opts *serverOptions) {
		opts.ShutdownTimeout = timeout
	}
}
