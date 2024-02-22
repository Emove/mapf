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

type nodeConfigRepo struct {
	*data.Pager
	data   *Data
	logger *log.Helper
}

func NewNodeConfigRepo(data *Data, logger log.Logger) (biz.NodeConfigRepo, error) {
	return &nodeConfigRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

// CreateNodeConfig 创建节点配置信息
func (repo *nodeConfigRepo) CreateNodeConfig(ctx context.Context, config *biz.NodeConfig) (*biz.NodeConfig, error) {
	err := repo.data.DB(ctx).Create(config).Error
	if err != nil && errors.Is(err, gorm.ErrDuplicatedKey) {
		return nil, dataerrors.NewDuplicatedKeyError("NodeConfig warehouse: [%s], node_id: [%d], config_item_id: [%d] Existed", config.WarehouseId, config.NodeRelationId, config.ConfigItemId)
	}
	return config, err
}

// UpdateNodeConfig 更新节点配置信息
func (repo *nodeConfigRepo) UpdateNodeConfig(ctx context.Context, config *biz.NodeConfig) (*biz.NodeConfig, error) {
	stmt := repo.data.DB(ctx).Model(&biz.NodeConfig{Model: &data.Model{ID: config.ID}})
	result := stmt.Updates(config.GetChanges())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, dataerrors.NewDuplicatedKeyError("NodeConfig warehouse: [%s], node_id: [%d], config_item_id: [%d] Existed", config.WarehouseId, config.NodeRelationId, config.ConfigItemId)
		}
		return nil, result.Error
	}
	return config, nil
}

// GetNodeConfigById 根据ID获取节点配置信息
func (repo *nodeConfigRepo) GetNodeConfigById(ctx context.Context, id int) (config *biz.NodeConfig, err error) {
	err = repo.data.DB(ctx).Where(biz.NodeConfig{Model: &data.Model{ID: id}}).First(&config).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// SelectNodeConfig 查询节点配置信息，可分页
func (repo *nodeConfigRepo) SelectNodeConfig(ctx context.Context, query *biz.NodeConfig, page *data.Page) ([]*biz.NodeConfig, *data.Page, error) {
	var configs []*biz.NodeConfig
	var err error
	db, page, err := repo.Pager.Prepare(repo.data.DB(ctx), query, page)
	if err != nil {
		return nil, nil, err
	}
	err = db.Find(&configs).Error
	return configs, page, err
}

// DeleteNodeConfigById 删除节点配置信息
func (repo *nodeConfigRepo) DeleteNodeConfigById(ctx context.Context, id int) error {
	return repo.data.DB(ctx).
		Model(&biz.NodeConfig{Model: &data.Model{ID: id}}).
		Updates(biz.NodeConfig{Model: &data.Model{IsDeleted: data.IsDeletedTrue}}).Error
}
