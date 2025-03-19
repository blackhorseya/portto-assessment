package httpx

import (
	"io"
	"net/http"
	"testing"
	"time"

	"log/slog"

	"github.com/gin-gonic/gin"
)

func TestGinServer_RunAndShutdown(t *testing.T) {
	// Create a test logger
	logger := slog.Default()

	// Create a new GinServer instance with debug mode enabled.
	server := NewGinServer(logger, true)

	// Register a test route.
	server.Router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// Start the server asynchronously using default options (host "localhost", port 8080).
	server.Run()

	// Wait briefly to allow the server to start.
	time.Sleep(100 * time.Millisecond)

	// Send a GET request to the /ping endpoint.
	resp, err := http.Get("http://localhost:8080/ping")
	if err != nil {
		t.Fatalf("failed to send GET request: %v", err)
	}
	defer resp.Body.Close()

	// Verify the status code is 200 OK.
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected status 200, got %d", resp.StatusCode)
	}

	// Read and verify the response body.
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read response body: %v", err)
	}
	if string(body) != "pong" {
		t.Fatalf("expected response body 'pong', got '%s'", string(body))
	}

	// Shutdown the server gracefully.
	err = server.Shutdown()
	if err != nil {
		t.Fatalf("failed to shutdown server: %v", err)
	}
}

func TestGinServer_ShutdownWithoutRun(t *testing.T) {
	// Create a test logger
	logger := slog.Default()

	// Create a new GinServer instance but do not run it.
	server := NewGinServer(logger, true)

	// Attempt to shutdown the server when it is not running.
	err := server.Shutdown()
	if err == nil {
		t.Fatalf("expected error 'server is not running', got nil")
	}

	expectedErr := "server is not running"
	if err.Error() != expectedErr {
		t.Fatalf("expected error '%s', got %v", expectedErr, err)
	}
}
