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

type nodeRelationRepo struct {
	*data.Pager
	data   *Data
	logger *log.Helper
}

func NewNodeRelationRepo(data *Data, logger log.Logger) (biz.NodeRelationRepo, error) {
	return &nodeRelationRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

// CreateNodeRelation 创建节点关联关系
func (repo *nodeRelationRepo) CreateNodeRelation(ctx context.Context, nodeRelation *biz.NodeRelation) (*biz.NodeRelation, error) {
	err := repo.data.DB(ctx).Create(nodeRelation).Error
	return nodeRelation, err
}

// GetNodeRelationById 根据ID获取节点关联关系
func (repo *nodeRelationRepo) GetNodeRelationById(ctx context.Context, id int) (relation *biz.NodeRelation, err error) {
	err = repo.data.DB(ctx).Where(biz.NodeRelation{Model: &data.Model{ID: id}}).First(&relation).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// UpdateNodeRelation 更新节点关联关系
func (repo *nodeRelationRepo) UpdateNodeRelation(ctx context.Context, relation *biz.NodeRelation) (*biz.NodeRelation, error) {
	stmt := repo.data.DB(ctx).Model(&biz.NodeRelation{Model: &data.Model{ID: relation.ID}})
	result := stmt.Updates(relation.GetChanges())
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return nil, dataerrors.NewDuplicatedKeyError("NodeRelation warehouse_id: [%d], node_id: [%d], affix_node_id: [%d] Existed", relation.WarehouseId, relation.NodeId, relation.AffixNodeId)
		}
		return nil, result.Error
	}
	return relation, nil
}

// SelectNodeRelation 查询节点关联关系。可分页
func (repo *nodeRelationRepo) SelectNodeRelation(ctx context.Context, query *biz.NodeRelation, page *data.Page) ([]*biz.NodeRelation, *data.Page, error) {
	var relations []*biz.NodeRelation
	var err error
	db, page, err := repo.Pager.Prepare(repo.data.DB(ctx), query, page)
	if err != nil {
		return nil, nil, err
	}
	err = db.Find(&relations).Error
	return relations, page, err
}

// DeleteNodeRelation 删除节点关联关系
func (repo *nodeRelationRepo) DeleteNodeRelation(ctx context.Context, id int) error {
	return repo.data.DB(ctx).
		Model(&biz.NodeRelation{Model: &data.Model{ID: id}}).
		Updates(biz.NodeRelation{Model: &data.Model{IsDeleted: data.IsDeletedTrue}}).Error
}
