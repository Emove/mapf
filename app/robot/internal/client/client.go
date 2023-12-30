package client

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/robot/internal/conf"
)

var ProviderSet = wire.NewSet(NewWarehouseClient)

func NewWarehouseClient(c *conf.Server, logger log.Logger) (v1.WarehouseClient, error) {
	conn, err := grpc.Dial(c.Warehouse, grpc.WithTransportCredentials(insecure.NewCredentials()))
	helper := log.NewHelper(logger)
	if err != nil {
		return nil, err
	}
	helper.Infow("[gRPC] connected to server", "warehouse", "addr", c.Warehouse)
	return v1.NewWarehouseClient(conn), err
}
