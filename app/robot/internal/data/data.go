package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/robot/internal/conf"
	"mapf/internal/data"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	ctx context.Context
	db  *gorm.DB
	rds *redis.Client
	whc v1.WarehouseClient
}

// NewData .
func NewData(serverConfig *conf.Server, dataConfig *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	db, err := data.NewPostgresDB(parse2database(dataConfig), helper)
	if err != nil {
		return nil, nil, err
	}
	rds, err := data.NewRedis(parse2redis(dataConfig), helper)
	if err != nil {
		return nil, nil, err
	}
	whc, err := newWarehouseClient(serverConfig, helper)
	if err != nil {
		return nil, nil, err
	}
	cleanup := func() {
		helper.Info("closing the data resources")
		_ = rds.Close()
	}
	return &Data{db: db, rds: rds, ctx: context.Background(), whc: whc}, cleanup, nil
}

func newWarehouseClient(c *conf.Server, logger *log.Helper) (v1.WarehouseClient, error) {
	conn, err := data.NewClientConnection(c.GetWarehouse(), c.GetGrpc().Timeout.AsDuration())
	if err != nil {
		return nil, err
	}
	logger.Infof("[gRPC] connected to server: warehouse, addr: %s", c.Warehouse)
	return v1.NewWarehouseClient(conn), err
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
