package data

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	dataerrors "mapf/internal/errors"
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

// CreateWarehouse 创建仓库
func (repo *warehouseRepo) CreateWarehouse(ctx context.Context, warehouse *biz.Warehouse) (*biz.Warehouse, error) {
	err := repo.data.DB(ctx).Create(warehouse).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return warehouse, dataerrors.NewDuplicatedKeyError("Warehouse Name: [%s] Existed", warehouse.Name)
	}
	return warehouse, err
}

// GetWarehouseByName 根据仓库名获取仓库信息
func (repo *warehouseRepo) GetWarehouseByName(ctx context.Context, name string) (warehouse *biz.Warehouse, err error) {
	err = repo.data.DB(ctx).Where(biz.Warehouse{Name: name}).First(&warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// GetWarehouseById 根据仓库ID获取仓库信息
func (repo *warehouseRepo) GetWarehouseById(ctx context.Context, id int) (warehouse *biz.Warehouse, err error) {
	err = repo.data.DB(ctx).Where(biz.Warehouse{Model: &data.Model{ID: id}}).First(&warehouse).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// UpdateWarehouseById 根据仓库ID更新仓库信息，仅当更新成功时返回True
func (repo *warehouseRepo) UpdateWarehouseById(ctx context.Context, id int, warehouse *biz.Warehouse) (bool, error) {
	stmt := repo.data.DB(ctx).
		Model(&biz.Warehouse{Model: &data.Model{ID: id}})
	values := map[string]interface{}{}
	if len(warehouse.Name) > 0 {
		values["name"] = warehouse.Name
	}
	if data.IsValidDefaultStatus(warehouse.IsDefault) {
		values["is_default"] = warehouse.IsDefault
	}
	if len(values) == 0 {
		return false, nil
	}
	result := stmt.Updates(values)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return false, dataerrors.NewDuplicatedKeyError("Warehouse Name: [%s] Existed", warehouse.Name)
		}
		return false, result.Error
	}
	return result.RowsAffected == 1, nil
}

// EnableWarehouseById 启动仓库，仅当更新成功时返回True
func (repo *warehouseRepo) EnableWarehouseById(ctx context.Context, id int) (bool, error) {
	result := repo.data.DB(ctx).
		Model(&biz.Warehouse{Model: &data.Model{ID: id}}).
		Where(biz.Warehouse{Status: data.DisableStatus}).
		Updates(biz.Warehouse{Status: data.EnableStatus})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected != 0, nil
}

// DisableWarehouseById 启动仓库，仅当更新成功时返回True
func (repo *warehouseRepo) DisableWarehouseById(ctx context.Context, id int) (bool, error) {
	result := repo.data.DB(ctx).
		Model(&biz.Warehouse{Model: &data.Model{ID: id}}).
		Where(biz.Warehouse{Status: data.EnableStatus, IsDefault: data.IsDefaultFalse}).
		Updates(biz.Warehouse{Status: data.DisableStatus})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected != 0, nil
}

// DeleteWarehouseById 删除仓库，仅当删除成功时返回True
func (repo *warehouseRepo) DeleteWarehouseById(ctx context.Context, id int) (bool, error) {
	result := repo.data.DB(ctx).
		Model(&biz.Warehouse{Model: &data.Model{ID: id}}).
		Where(biz.Warehouse{Model: &data.Model{IsDeleted: data.IsDeletedFalse}, IsDefault: data.IsDefaultFalse}).
		Updates(biz.Warehouse{Model: &data.Model{IsDeleted: data.IsDeletedTrue}})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected != 0, nil
}
