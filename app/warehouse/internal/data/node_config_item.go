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
	*data.Pager
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
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, dataerrors.NewDuplicatedKeyError("NodeConfigItem Code: [%s] Existed", item.Code)
	}
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

// GetNodeConfigItemById 根据ID查询节点配置项
func (repo *nodeConfigItemRepo) GetNodeConfigItemById(ctx context.Context, id int) (item *biz.NodeConfigItem, err error) {
	err = repo.data.DB(ctx).Where(biz.NodeConfigItem{Model: &data.Model{ID: id}}).First(&item).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
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

// SelectNodeConfigItem 查询节点配置项，可分页
func (repo *nodeConfigItemRepo) SelectNodeConfigItem(ctx context.Context, query *biz.NodeConfigItem, page *data.Page) ([]*biz.NodeConfigItem, *data.Page, error) {
	var items []*biz.NodeConfigItem
	var err error
	db, page, err := repo.Pager.Prepare(repo.data.DB(ctx), query, page)
	if err != nil {
		return nil, nil, err
	}
	err = db.Find(&items).Error
	return items, page, err
}
