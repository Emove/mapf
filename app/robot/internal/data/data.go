package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
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
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	helper := log.NewHelper(logger)
	database := c.GetDatabase()
	db, err := data.NewPostgresDB(&data.Database{
		Host:         database.GetHost(),
		Port:         database.GetPort(),
		User:         database.GetUser(),
		Password:     database.GetPassword(),
		DatabaseName: database.GetDatabase(),
	}, helper)
	if err != nil {
		return nil, nil, err
	}
	rdsConfig := c.GetRedis()
	rds, err := data.NewRedis(&data.Redis{
		Host:         rdsConfig.GetHost(),
		Port:         rdsConfig.GetPort(),
		Password:     rdsConfig.GetPassword(),
		Db:           rdsConfig.GetDb(),
		ReadTimeout:  rdsConfig.ReadTimeout.AsDuration(),
		WriteTimeout: rdsConfig.WriteTimeout.AsDuration(),
	}, helper)
	cleanup := func() {
		helper.Info("closing the data resources")
		_ = rds.Close()
	}
	return &Data{db: db, rds: rds, ctx: context.Background()}, cleanup, nil
}
