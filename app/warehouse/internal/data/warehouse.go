package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
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

func (repo *warehouseRepo) GetWarehouseByName(ctx context.Context, name string) (warehouse *biz.Warehouse, err error) {
	err = repo.data.DB(ctx).Where(&biz.Warehouse{Name: name}).First(&warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return warehouse, err
}

func (repo *warehouseRepo) GetWarehouseById(ctx context.Context, id int) (warehouse *biz.Warehouse, err error) {
	err = repo.data.DB(ctx).Where(biz.Warehouse{Model: &data.Model{ID: id}}).First(&warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return warehouse, err
}
