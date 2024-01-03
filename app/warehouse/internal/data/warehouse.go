package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

var _ biz.WarehouseRepo = (*warehouseRepo)(nil)

type warehouseRepo struct {
	data   *Data
	logger *log.Helper
}

func NewWarehouseRepo(data *Data, logger log.Logger) (biz.WarehouseRepo, error) {
	return &warehouseRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

func (repo *warehouseRepo) CreateWarehouse(ctx context.Context, warehouse *biz.Warehouse) (*biz.Warehouse, error) {
	err := repo.data.DB(ctx).Create(warehouse).Error
	return warehouse, err
}
