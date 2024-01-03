package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	robot "mapf/api/robot/v1"
	warehouse "mapf/api/warehouse/v1"
	"mapf/app/plan/internal/conf"
	"mapf/internal/data"
	"strings"
	"time"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	ctx context.Context
	db  *gorm.DB
	rds *redis.Client
	whc warehouse.WarehouseServiceClient
	rbc robot.RobotServiceClient
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	instance := &Data{}
	var err error
	instance.db, err = data.NewPostgresDB(parse2database(c), logger, helper)
	if err != nil {
		return nil, nil, err
	}
	instance.rds, err = data.NewRedis(parse2redis(c), helper)
	if err != nil {
		return nil, nil, err
	}
	for _, service := range c.GetServices() {
		serviceName, addr, timeout := strings.ToLower(service.Name), service.Addr, service.Timeout.AsDuration()
		if serviceName == "warehouse" {
			instance.whc, err = newWarehouseClient(addr, timeout)
		} else if serviceName == "robot" {
			instance.rbc, err = newRobotClient(addr, timeout)
		}
		if err != nil {
			return nil, nil, err
		}
		helper.Infof("[gRPC] connected to server: %s, addr: %s, timeout: %ds", service.Name, service.Addr, service.Timeout.Seconds)
	}
	cleanup := func() {
		helper.Info("closing the data resources")
		_ = instance.rds.Close()
	}
	return instance, cleanup, nil
}

func newWarehouseClient(addr string, timeout time.Duration) (warehouse.WarehouseServiceClient, error) {
	conn, err := data.NewClientConnection(addr, timeout)
	return warehouse.NewWarehouseServiceClient(conn), err
}

func newRobotClient(addr string, timeout time.Duration) (robot.RobotServiceClient, error) {
	conn, err := data.NewClientConnection(addr, timeout)
	return robot.NewRobotServiceClient(conn), err
}

func parse2database(dataConfig *conf.Data) *data.Database {
	database := dataConfig.GetDatabase()
	return &data.Database{
		Host:         database.GetHost(),
		Port:         database.GetPort(),
		User:         database.GetUser(),
		Password:     database.GetPassword(),
		DatabaseName: database.GetDatabase(),
	}
}

func parse2redis(dataConfig *conf.Data) *data.Redis {
	rds := dataConfig.GetRedis()
	return &data.Redis{
		Host:         rds.GetHost(),
		Port:         rds.GetPort(),
		Password:     rds.GetPassword(),
		Db:           rds.GetDb(),
		ReadTimeout:  rds.ReadTimeout.AsDuration(),
		WriteTimeout: rds.WriteTimeout.AsDuration(),
	}
}
