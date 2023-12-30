package client

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	robot "mapf/api/robot/v1"
	warehouse "mapf/api/warehouse/v1"
	"mapf/app/plan/internal/conf"
)

var ProviderSet = wire.NewSet(NewWarehouseClient, NewRobotClient)

func NewWarehouseClient(c *conf.Server, logger log.Logger) (warehouse.WarehouseClient, error) {
	conn, err := grpc.Dial(c.Warehouse, grpc.WithTransportCredentials(insecure.NewCredentials()))
	helper := log.NewHelper(logger)
	if err != nil {
		return nil, err
	}
	helper.Infof("[grpc] do connect to [warehouse:%s] server succeed", c.Warehouse)
	return warehouse.NewWarehouseClient(conn), err
}

func NewRobotClient(c *conf.Server, logger log.Logger) (robot.RobotClient, error) {
	conn, err := grpc.Dial(c.Robot, grpc.WithTransportCredentials(insecure.NewCredentials()))
	helper := log.NewHelper(logger)
	if err != nil {
		return nil, err
	}
	helper.Infof("[grpc] do connect to [robot:%s] server success", c.Robot)
	return robot.RobotClient(conn), err
}
