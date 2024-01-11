package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/internal/data"
	"mapf/internal/data/tx"
	data_errors "mapf/internal/errors"
)

type Warehouse struct {
	*data.Model
	Name      string `gorm:"uniqueIndex:uidx_name;type:varchar(64);not null'"`
	IsDefault int    `gorm:"type:smallint;default:2;not null"`
	Status    int    `gorm:"type:smallint;default:1;not null"`
}

func (*Warehouse) TableName() string {
	return "warehouse"
}

type WarehouseRepo interface {
	// CreateWarehouse 创建仓库
	CreateWarehouse(context.Context, *Warehouse) (*Warehouse, error)
	// GetWarehouseByName 根据仓库名获取仓库信息
	GetWarehouseByName(context.Context, string) (*Warehouse, error)
	// GetWarehouseById 根据仓库ID获取仓库信息
	GetWarehouseById(context.Context, int) (*Warehouse, error)
	// UpdateWarehouseById 根据仓库ID更新仓库信息，仅当更新成功时返回True
	UpdateWarehouseById(context.Context, int, *Warehouse) (bool, error)
	// EnableWarehouseById 启动仓库，仅当更新成功时返回True
	EnableWarehouseById(context.Context, int) (bool, error)
	// DisableWarehouseById 启动仓库，仅当更新成功时返回True
	DisableWarehouseById(context.Context, int) (bool, error)
	// DeleteWarehouseById 删除仓库，仅当删除成功时返回True
	DeleteWarehouseById(context.Context, int) (bool, error)
}

type WarehouseUsecase struct {
	tm     tx.Transaction
	repo   WarehouseRepo
	logger *log.Helper
}

func NewWarehouseUsecase(tm tx.Transaction, repo WarehouseRepo, logger log.Logger) *WarehouseUsecase {
	return &WarehouseUsecase{tm: tm, repo: repo, logger: log.NewHelper(logger)}
}

// CreateWarehouse 创建仓库
func (uc *WarehouseUsecase) CreateWarehouse(ctx context.Context, warehouse *Warehouse) (*Warehouse, error) {
	return warehouse, uc.tm.InTx(ctx, func(ctx context.Context) error {
		var err error
		if warehouse, err = uc.repo.CreateWarehouse(ctx, warehouse); err != nil {
			return err
		}
		return nil
	})
}

// GetWarehouseByName 根据仓库名获取仓库信息
func (uc *WarehouseUsecase) GetWarehouseByName(ctx context.Context, name string) (*Warehouse, error) {
	return uc.repo.GetWarehouseByName(ctx, name)
}

// GetWarehouseById 根据仓库ID获取仓库信息
func (uc *WarehouseUsecase) GetWarehouseById(ctx context.Context, id int) (*Warehouse, error) {
	return uc.repo.GetWarehouseById(ctx, id)
}

// UpdateWarehouseById 根据仓库ID更新仓库信息
func (uc *WarehouseUsecase) UpdateWarehouseById(ctx context.Context, id int, warehouse *Warehouse) error {
	_, err := uc.repo.UpdateWarehouseById(ctx, id, warehouse)
	return err
}

// EnableWarehouseById 根据ID启动仓库
func (uc *WarehouseUsecase) EnableWarehouseById(ctx context.Context, id int) error {
	succeed, err := uc.repo.EnableWarehouseById(ctx, id)
	if err != nil {
		return err
	}
	if succeed {
		// 发布仓库启用事件
	}
	return nil
}

// DisableWarehouseById 根据仓库ID禁用仓库
func (uc *WarehouseUsecase) DisableWarehouseById(ctx context.Context, id int) error {
	warehouse, err := uc.getWarehouseById(ctx, id)
	if err != nil {
		return err
	}
	if warehouse.IsDefault == data.IsDefaultTrue {
		return errors.Forbidden("DEFAULT_RESOURCE", "disable a default warehouse is not allowed")
	}
	succeed, err := uc.repo.DisableWarehouseById(ctx, id)
	if err != nil {
		return err
	}
	if succeed {
		// 发布仓库禁用事件
	}
	return nil
}

func (uc *WarehouseUsecase) DeleteWarehouseById(ctx context.Context, id int) error {
	warehouse, err := uc.getWarehouseById(ctx, id)
	if err != nil {
		return err
	}
	if warehouse.IsDefault == data.IsDefaultTrue {
		return data_errors.NewDefaultResourceForbiddenError("delete a default Warehouse is not allowed")
	}
	succeed, err := uc.repo.DeleteWarehouseById(ctx, id)
	if err != nil {
		return err
	}
	if succeed {
		// 发布仓库删除事件
	}
	return nil
}

func (uc *WarehouseUsecase) getWarehouseById(ctx context.Context, id int) (*Warehouse, error) {
	warehouse, err := uc.repo.GetWarehouseById(ctx, id)
	if err != nil {
		return nil, err
	}
	if warehouse == nil {
		return nil, data_errors.NewResourceNotFoundError("warehouse id: %d not found", id)
	}
	return warehouse, nil
}
