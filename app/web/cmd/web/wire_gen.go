// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/web/internal/biz"
	"mapf/app/web/internal/conf"
	"mapf/app/web/internal/data"
	"mapf/app/web/internal/server"
	"mapf/app/web/internal/service"
)

import (
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, logger log.Logger) (*kratos.App, func(), error) {
	dataData, cleanup, err := data.NewData(confData, logger)
	if err != nil {
		return nil, nil, err
	}
	warehouseUsecase := biz.NewWarehouseUsecase(dataData, logger)
	webService := service.NewWebService(warehouseUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, webService, logger)
	httpServer := server.NewHTTPServer(confServer, webService, logger)
	app := newApp(logger, grpcServer, httpServer)
	return app, func() {
		cleanup()
	}, nil
}
