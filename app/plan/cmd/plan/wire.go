//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"mapf/app/plan/internal/biz"
	"mapf/app/plan/internal/client"
	"mapf/app/plan/internal/conf"
	"mapf/app/plan/internal/data"
	"mapf/app/plan/internal/server"
	"mapf/app/plan/internal/service"

	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, client.ProviderSet, newApp))
}
