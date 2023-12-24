package data

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"mapf/internal/conf"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
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
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	helper := log.NewHelper(logger)
	db, err := newDB(c, helper)
	if err != nil {
		return nil, nil, err
	}
	rds, err := newRedis(c, helper)
	if err != nil {
		return nil, nil, err
	}
	data := &Data{
		ctx: context.Background(),
		db:  db,
		rds: rds,
	}
	return data, cleanup, nil
}

func newDB(c *conf.Data, logger *log.Helper) (*gorm.DB, error) {
	database := c.GetDatabase()
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		database.GetHost(), database.GetUser(), database.GetPassword(), database.GetDatabase(), database.GetPort())
	logger.Infow("database dsn", dsn)
	return gorm.Open(postgres.Open(dsn), &gorm.Config{})
}

func newRedis(c *conf.Data, logger *log.Helper) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:         c.Redis.GetAddr(),
		Password:     c.Redis.GetPassword(), // no password set
		DB:           0,                     // use default DB
		ReadTimeout:  c.Redis.GetReadTimeout().AsDuration(),
		WriteTimeout: c.Redis.GetWriteTimeout().AsDuration(),
	})

	logger.Infow("redis configuration", c.Redis.String())
	resp := rds.Ping(context.Background())
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return rds, nil
}
