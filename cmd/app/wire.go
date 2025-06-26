//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"context"

	"thmanyah/internal/conf"
	"thmanyah/internal/modules/cms"
	"thmanyah/internal/modules/discover"
	"thmanyah/internal/postgres"
	"thmanyah/internal/server"
	"thmanyah/keys"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

func wireApp(context.Context, log.Logger, *conf.Server, *conf.Data) (*kratos.App, error) {
	panic(
		wire.Build(
			postgres.NewPgPool,
			keys.Provider,
			server.ProviderSet,
			cms.ProviderSet,
			discover.ProviderSet,
			newApp,
		),
	)
}
