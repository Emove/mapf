package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"gorm.io/gorm"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
	data_errors "mapf/internal/errors"
)

var _ biz.NodeTypeRepo = (*nodeTypeRepo)(nil)

type nodeTypeRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeTypeRepo(data *Data, logger log.Logger) (biz.NodeTypeRepo, error) {
	return &nodeTypeRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}

// CreateNodeType 在创建NodeType
func (repo *nodeTypeRepo) CreateNodeType(ctx context.Context, nodeType *biz.NodeType) (*biz.NodeType, error) {
	err := repo.data.DB(ctx).Create(nodeType).Error
	return nodeType, err
}

// BatchCreateNodeTypes 批量创建NodeType
func (repo *nodeTypeRepo) BatchCreateNodeTypes(ctx context.Context, nodeTypes []*biz.NodeType) ([]*biz.NodeType, error) {
	db := repo.data.DB(ctx)
	for i := 0; i < len(nodeTypes); i++ {
		if err := db.Create(nodeTypes[i]).Error; err != nil {
			return nodeTypes, err
		}
	}
	return nodeTypes, nil
}

// GetNodeTypeById 根据节点类型ID获取节点类型
func (repo *nodeTypeRepo) GetNodeTypeById(ctx context.Context, id int) (nodeType *biz.NodeType, err error) {
	err = repo.data.DB(ctx).Where(biz.NodeType{Model: &data.Model{ID: id}}).First(&nodeType).Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		return nil, nil
	}
	return
}

// GetNodeTypesByIds 根据节点类型ID列表获取节点类型列表
func (repo *nodeTypeRepo) GetNodeTypesByIds(ctx context.Context, ids []int) (nodeTypes []*biz.NodeType, err error) {
	err = repo.data.DB(ctx).Where(ids).Find(&nodeTypes).Error
	return
}

// UpdateNodeTypeById 根据节点类型ID更新节点类型
func (repo *nodeTypeRepo) UpdateNodeTypeById(ctx context.Context, id int, nodeType *biz.NodeType) (bool, error) {
	stmt := repo.data.DB(ctx).
		Model(&biz.NodeType{Model: &data.Model{ID: id}})
	values := make(map[string]interface{})
	if len(nodeType.Code) > 0 {
		values["code"] = nodeType.Code
	}
	if len(nodeType.Name) > 0 {
		values["name"] = nodeType.Name
	}
	if data.IsValidDefaultStatus(nodeType.IsDefault) {
		values["is_default"] = nodeType.IsDefault
	}
	if len(values) == 0 {
		return false, nil
	}
	result := stmt.Updates(values)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrDuplicatedKey) {
			return false, data_errors.NewDuplicatedKeyError("NodeType Code: [%s] Existed", nodeType.Code)
		}
		return false, result.Error
	}
	return result.RowsAffected == 1, nil
}

// DeleteNodeTypeById 根据节点类型ID删除节点类型，仅当删除成功时返回True
func (repo *nodeTypeRepo) DeleteNodeTypeById(ctx context.Context, id int) (bool, error) {
	result := repo.data.DB(ctx).
		Model(&biz.NodeType{Model: &data.Model{ID: id}}).
		Where(biz.NodeType{Model: &data.Model{IsDeleted: data.IsDeletedFalse}, IsDefault: data.IsDefaultFalse}).
		Updates(biz.NodeType{Model: &data.Model{IsDeleted: data.IsDeletedTrue}})
	if result.Error != nil {
		return false, result.Error
	}
	return result.RowsAffected != 0, nil
}
