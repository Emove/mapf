package data

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/app/warehouse/internal/conf"
	"mapf/internal/data"
	"mapf/internal/data/tx"
	"mapf/internal/event"
	"net/url"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewTransaction, NewEventCenterConfiguration, NewWarehouseRepo,
	NewNodeRepo, NewNodeConfigRepo, NewNodeConfigItemRepo, NewNodeTypeRepo, NewAffixNodeRepo, NewNodeRelationRepo,
	NewNodeTagRepo, NewNodeDiagramRepo)

var models = []interface{}{&biz.Warehouse{}, &biz.Node{}, &biz.NodeConfig{}, &biz.NodeConfigItem{}, &biz.NodeType{},
	&biz.AffixNode{}, &biz.NodeRelation{}, &biz.NodeTag{}, &biz.NodeDiagram{}}

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
		LogLevel:     database.LogLevel,
	}, logger, helper)
	if err != nil {
		return nil, nil, err
	}
	if err = autoMigration(db); err != nil {
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

func NewTransaction(data *Data) tx.Transaction {
	return data
}

func NewEventCenterConfiguration(c *conf.Data, logger log.Logger) (event.Configuration, func(), error) {
	config := c.Event
	parse, err := url.Parse(c.Event.Addr)
	if err != nil {
		return nil, nil, err
	}
	if parse.Scheme == event.AmqpScheme {
		return event.NewRabbitMqConfiguration(config.Addr, logger)
	}
	return nil, nil, errors.New(fmt.Sprintf("event center type not supported, addr: %s", config.Addr))
}

func (d *Data) DB(ctx context.Context) *gorm.DB {
	if db := tx.GetTxFromContext(ctx); db != nil {
		return db
	}
	return d.db
}

func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(db *gorm.DB) error {
		return fn(tx.SetTxToContext(ctx, db))
	})
}

func autoMigration(db *gorm.DB) error {
	var err error
	for _, model := range models {
		if err = db.AutoMigrate(model); err != nil {
			return err
		}
	}
	return nil
}
