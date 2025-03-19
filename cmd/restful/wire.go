//go:build wireinject

//go:generate wire

package restful

import (
	"portto/internal/handler"
	"portto/internal/repository"
	"portto/internal/shared/configx"
	"portto/internal/shared/pgx"
	"portto/pkg/otelx"

	"portto/pkg/contextx"
	"portto/pkg/httpx"

	"github.com/google/wire"
)

func newGinServer(ctx contextx.Contextx, appConfig *configx.Application) *httpx.GinServer {
	return httpx.NewGinServer(ctx.Logger, appConfig.Verbose)
}

func newOTelSDK(appConfig *configx.Application) (*otelx.SDK, func(), error) {
	return otelx.SetupSDK(appConfig.OTel.Target, "portto", appConfig.Verbose)
}

func NewServer(ctx contextx.Contextx, appConfig *configx.Application) (*Server, func(), error) {
	panic(wire.Build(
		wire.Struct(new(Server), "*"),
		newGinServer,
		newOTelSDK,
		handler.RegisterRoutes,
		repository.NewCoinRepository,
		pgx.NewClient,
	))
}
