//go:build wireinject

//go:generate wire

package restful

import (
	"portto/internal/handler"
	"portto/internal/repository"
	"portto/internal/shared/configx"
	"portto/internal/shared/pgx"

	"github.com/blackhorseya/go-libs/contextx"
	"github.com/blackhorseya/go-libs/httpx"
	"github.com/google/wire"
)

func newGinServer(ctx contextx.Contextx, appConfig *configx.Application) *httpx.GinServer {
	return httpx.NewGinServer(ctx.Logger, appConfig.Verbose)
}

func NewServer(ctx contextx.Contextx, appConfig *configx.Application) (*Server, func(), error) {
	panic(wire.Build(
		wire.Struct(new(Server), "*"),
		newGinServer,
		handler.RegisterRoutes,
		repository.NewCoinRepository,
		pgx.NewClient,
	))
}
