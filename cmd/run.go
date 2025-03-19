package cmd

import (
	"log/slog"
	"os"
	"os/signal"
	"portto/cmd/restful"
	"portto/internal/shared/configx"
	"syscall"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"portto/pkg/contextx"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run the server",
	Run: func(cmd *cobra.Command, args []string) {
		var appConfig configx.Application
		err := viper.Unmarshal(&appConfig)
		if err != nil {
			slog.ErrorContext(cmd.Context(), "Unable to unmarshal config", "error", err)
			return
		}

		err = appConfig.SetupLogger()
		if err != nil {
			slog.ErrorContext(cmd.Context(), "Unable to setup logger", "error", err)
			return
		}

		// extended context
		ctx := contextx.WithContext(cmd.Context())
		ctx.Debug("Config loaded", "config", appConfig)

		server, clean, err := restful.NewServer(ctx, &appConfig)
		if err != nil {
			ctx.Error("Failed to create server", "error", err)
			return
		}
		defer clean()

		err = server.Start(ctx)
		if err != nil {
			ctx.Error("Failed to start server", "error", err)
			return
		}
		ctx.Info("Server started", "host", appConfig.Host, "port", appConfig.Port)

		signalChan := make(chan os.Signal, 1)
		signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM)
		<-signalChan

		ctx.Info("Shutting down server...")
		err = server.Stop(ctx)
		if err != nil {
			ctx.Error("Server shutdown failed", "error", err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.
	runCmd.PersistentFlags().String("host", "localhost", "Host to run the server on")
	_ = viper.BindPFlag("host", runCmd.PersistentFlags().Lookup("host"))

	runCmd.PersistentFlags().Int("port", 8080, "Port to run the server on")
	_ = viper.BindPFlag("port", runCmd.PersistentFlags().Lookup("port"))
}
