package biz

import (
	"context"
	"errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	v1 "mapf/api/warehouse/v1"
	"mapf/internal/data"
	"mapf/internal/data/tx"
)

type Warehouse struct {
	*data.Model
	Name      string `gorm:"uniqueIndex:uidx_name;type:varchar(64);not null'"`
	IsDefault int    `gorm:"type:smallint;default:0;not null"`
	Status    int    `gorm:"type:smallint;default:1;not null"`
}

func (*Warehouse) TableName() string {
	return "warehouse"
}

type WarehouseRepo interface {
	CreateWarehouse(ctx context.Context, warehouse *Warehouse) (*Warehouse, error)
	GetWarehouseByName(ctx context.Context, name string) (*Warehouse, error)
	GetWarehouseById(ctx context.Context, id int) (*Warehouse, error)
}

type WarehouseUsecase struct {
	tm           tx.Transaction
	repo         WarehouseRepo
	nodeTypeRepo NodeTypeRepo
	logger       *log.Helper
}

func NewWarehouseUsecase(tm tx.Transaction, repo WarehouseRepo, nodeTypeRepo NodeTypeRepo, logger log.Logger) *WarehouseUsecase {
	return &WarehouseUsecase{tm: tm, repo: repo, nodeTypeRepo: nodeTypeRepo, logger: log.NewHelper(logger)}
}

// CreateWarehouse 创建仓库
func (uc *WarehouseUsecase) CreateWarehouse(ctx context.Context, warehouse *Warehouse, nodeTypes []*NodeType) (*Warehouse, []*NodeType, error) {
	return warehouse, nodeTypes, uc.tm.InTx(ctx, func(ctx context.Context) error {
		var err error
		if warehouse, err = uc.repo.CreateWarehouse(ctx, warehouse); err != nil {
			if errors.Is(err, gorm.ErrDuplicatedKey) {
				return v1.ErrorWarehouseExisted("warehouse [%s] existed", warehouse.Name)
			}
			return err
		}
		global := &NodeType{WarehouseId: warehouse.ID, Code: "GLOBAL", Name: "默认节点", IsDefault: data.IsDefaultTrue}
		if _, err = uc.nodeTypeRepo.CreateNodeType(ctx, global); err != nil {
			return err
		}

		if len(nodeTypes) > 0 {
			for _, nodeType := range nodeTypes {
				nodeType.WarehouseId = warehouse.ID
			}
			if _, err = uc.nodeTypeRepo.BatchCreateNodeTypes(ctx, nodeTypes); err != nil {
				if errors.Is(err, gorm.ErrDuplicatedKey) {
					return v1.ErrorNodeTypeExisted("Duplicated NodeType in Warehouse [%s]", warehouse.Name)
				}
				return err
			}
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
