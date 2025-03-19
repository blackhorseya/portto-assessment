package loggerx

import (
	"testing"
)

func TestZapLogger(t *testing.T) {
	logger, err := NewZapLogger(Options{
		Level:  "debug",
		Format: "json",
	})
	if err != nil {
		t.Errorf("failed to init logger: %v", err)
	}

	// Test the logger
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
}

func TestSlogLogger(t *testing.T) {
	logger, err := NewSlogLogger(WithFormat("json"), WithLevel("debug"))
	if err != nil {
		t.Errorf("failed to init logger: %v", err)
	}

	// Test the logger
	logger.Debug("debug message")
	logger.Info("info message")
	logger.Warn("warn message")
	logger.Error("error message")
}
