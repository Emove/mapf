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

type nodeConfigItemRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeConfigItemRepo(data *Data, logger log.Logger) (biz.NodeConfigItemRepo, error) {
	return &nodeConfigItemRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

// CreateNodeConfigItem 创建仓库节点配置项
func (repo *nodeConfigItemRepo) CreateNodeConfigItem(ctx context.Context, item *biz.NodeConfigItem) (*biz.NodeConfigItem, error) {
	err := repo.data.DB(ctx).Create(item).Error
	return item, err
}

// UpdateNodeConfigItem 更新节点配置项
func (repo *nodeConfigItemRepo) UpdateNodeConfigItem(ctx context.Context, item *biz.NodeConfigItem) (*biz.NodeConfigItem, error) {
	stmt := repo.data.DB(ctx).Model(&biz.NodeConfigItem{Model: &data.Model{ID: item.ID}})
	result := stmt.Updates(item.GetChanges())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, dataerrors.NewDuplicatedKeyError("NodeConfigItem Code: [%s] Existed", item.Code)
		}
		return nil, result.Error
	}
	return item, nil
}

// GetNodeConfigItemByWarehouseIdAndNodeTypeId 查询仓库节点配置项，节点类型可为空
func (repo *nodeConfigItemRepo) GetNodeConfigItemByWarehouseIdAndNodeTypeId(ctx context.Context, warehouseId int, nodeTypeId *int) (items []*biz.NodeConfigItem, err error) {
	query := biz.NodeConfigItem{WarehouseId: warehouseId}
	if nodeTypeId != nil {
		query.NodeTypeId = *nodeTypeId
	}
	err = repo.data.DB(ctx).Where(query).Find(&items).Error
	return
}
