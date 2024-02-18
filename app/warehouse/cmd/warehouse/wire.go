//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mapf/app/warehouse/internal/biz"
	"mapf/app/warehouse/internal/conf"
	"mapf/app/warehouse/internal/data"
	"mapf/app/warehouse/internal/server"
	"mapf/app/warehouse/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(
		server.ProviderSet,
		data.ProviderSet,
		//publisher.ProviderSet,
		//subscriber.ProviderSet,
		biz.ProviderSet,
		service.ProviderSet,
		newApp),
	)
}
