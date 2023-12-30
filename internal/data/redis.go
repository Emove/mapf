package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/redis/go-redis/v9"
	"strings"
	"time"
)

type Redis struct {
	Host         string
	Port         string
	Password     string
	Db           int32
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
}

func NewRedis(config *Redis, logger *log.Helper) (*redis.Client, error) {
	rds := redis.NewClient(&redis.Options{
		Addr:         strings.Join([]string{config.Host, config.Port}, ":"),
		Password:     config.Password,
		DB:           int(config.Db),
		ReadTimeout:  config.ReadTimeout,
		WriteTimeout: config.WriteTimeout,
	})

	logger.Infow("redis configuration", config)
	resp := rds.Ping(context.Background())
	if resp.Err() != nil {
		return nil, resp.Err()
	}
	return rds, nil
}
