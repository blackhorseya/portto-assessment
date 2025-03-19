//go:build external

package restful

import (
	"context"
	"os"
	"os/signal"
	"portto/internal/shared/configx"
	"syscall"
	"testing"

	"github.com/blackhorseya/go-libs/contextx"
	_ "github.com/joho/godotenv/autoload"
	"github.com/spf13/viper"
)

func TestServerRun(t *testing.T) {
	var appConfig configx.Application
	err := viper.Unmarshal(&appConfig)
	if err != nil {
		t.Fatalf("Unable to unmarshal config: %v", err)
	}
	appConfig.Port = 8080

	err = appConfig.SetupLogger()
	if err != nil {
		t.Fatalf("Unable to setup logger: %v", err)
	}

	ctx := contextx.WithContext(context.Background())
	server, clean, err := NewServer(ctx, &appConfig)
	if err != nil {
		t.Fatalf("Failed to create server: %v", err)
	}
	defer clean()

	err = server.Start(ctx)
	if err != nil {
		t.Fatalf("Failed to start server: %v", err)
	}

	t.Logf("Server started on %s:%d", appConfig.Host, appConfig.Port)

	signalChan := make(chan os.Signal, 1)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
	<-signalChan

	err = server.Stop(ctx)
	if err != nil {
		t.Fatalf("Server shutdown failed: %v", err)
	}
}
