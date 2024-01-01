// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
	"mapf/app/warehouse/internal/conf"
	"mapf/app/warehouse/internal/data"
	"mapf/app/warehouse/internal/server"
	"mapf/app/warehouse/internal/service"
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
	warehouseRepo, err := data.NewWarehouseRepo(dataData, logger)
	if err != nil {
		cleanup()
		return nil, nil, err
	}
	warehouseUsecase := biz.NewWarehouseUsecase(warehouseRepo, logger)
	warehouseService := service.NewWarehouseService(warehouseUsecase)
	grpcServer := server.NewGRPCServer(confServer, warehouseService, logger)
	app := newApp(logger, grpcServer)
	return app, func() {
		cleanup()
	}, nil
}
