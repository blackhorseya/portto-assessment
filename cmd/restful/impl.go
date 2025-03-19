package restful

import (
	"portto/internal/shared/configx"
	"portto/pkg/otelx"
	"time"

	"portto/pkg/contextx"
	"portto/pkg/httpx"
)

// @title Swagger Example API
// @version 1.0
// @BasePath /api

type Server struct {
	appConfig  *configx.Application
	ginServer  *httpx.GinServer
	initRouter httpx.InitRouterFn
	otelSDK    *otelx.SDK
}

func (x *Server) Start(ctx contextx.Contextx) error {
	if x.ginServer != nil {
		// Initialize the Gin server with the provided router function.
		x.initRouter(x.ginServer.Router)

		x.ginServer.Run(httpx.WithHost(x.appConfig.Host), httpx.WithPort(x.appConfig.Port), httpx.WithLogger(ctx.Logger))
	}

	return nil
}

func (x *Server) Stop(ctx contextx.Contextx) error {
	if x.ginServer != nil {
		return x.ginServer.Shutdown(httpx.WithShutdownTimeout(10*time.Second), httpx.WithLogger(ctx.Logger))
	}

	return nil
}
