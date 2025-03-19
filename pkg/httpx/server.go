package httpx

import (
	"context"
	"errors"
	"fmt"
	"log/slog"
	"net/http"
	"portto/pkg/loggerx"
	"time"

	"github.com/gin-gonic/gin"
)

type InitRouterFn func(*gin.Engine)

// GinServer is a wrapper around http.Server and gin.Engine for managing routes and server lifecycle.
type GinServer struct {
	httpserver *http.Server
	Router     *gin.Engine
}

// NewGinServer creates a new GinServer
func NewGinServer(log *slog.Logger, debug bool) *GinServer {
	middlewares := []gin.HandlerFunc{
		loggerx.GinTraceLoggingMiddleware(log),
	}

	gin.SetMode(gin.ReleaseMode)
	if debug {
		gin.SetMode(gin.DebugMode)
	}

	router := gin.New()
	router.Use(middlewares...)

	return &GinServer{
		httpserver: nil,
		Router:     router,
	}
}

// Run starts the HTTP server asynchronously using the Options pattern.
// Default configuration is: address ":8080", default logger, and shutdown timeout of 5 seconds.
// Options such as WithAddr, WithLogger, and WithShutdownTimeout can be used to override the defaults.
func (s *GinServer) Run(opts ...ServerOption) {
	// Set default options.
	options := &serverOptions{
		Host:            "localhost",
		Port:            8080,
		Logger:          slog.Default(),
		ShutdownTimeout: 5 * time.Second,
	}
	// Override defaults with provided options.
	for _, opt := range opts {
		opt(options)
	}

	// Construct the address string.
	addr := fmt.Sprintf("%s:%d", options.Host, options.Port)

	// Initialize http.Server based on the configured options.
	s.httpserver = &http.Server{
		Addr:              addr,
		Handler:           s.Router,
		ReadHeaderTimeout: time.Second,
	}

	// Start the server asynchronously.
	go func() {
		if err := s.httpserver.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			options.Logger.Error("failed to start gin server", "error", err)
		}
	}()
}

// Shutdown gracefully shuts down the server using the configured shutdown timeout.
func (s *GinServer) Shutdown(opts ...ServerOption) error {
	if s.httpserver == nil {
		return errors.New("server is not running")
	}

	// Set default shutdown options.
	options := &serverOptions{
		ShutdownTimeout: 5 * time.Second,
		Logger:          slog.Default(),
	}
	// Override defaults with provided options.
	for _, opt := range opts {
		opt(options)
	}

	ctx, cancel := context.WithTimeout(context.Background(), options.ShutdownTimeout)
	defer cancel()
	return s.httpserver.Shutdown(ctx)
}
