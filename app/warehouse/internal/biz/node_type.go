package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/internal/data"
	"mapf/internal/data/tx"
	data_errors "mapf/internal/errors"
)

type NodeType struct {
	*data.Model
	Code      string `gorm:"uniqueIndex:uidx_code;type:varchar(63);not null"`
	Name      string `gorm:"type:varchar(63);not null"`
	IsDefault int    `gorm:"type:smallint;default:2;not null"`
}

func (*NodeType) TableName() string {
	return "node_type"
}

type NodeTypeRepo interface {
	// CreateNodeType 在创建NodeType
	CreateNodeType(context.Context, *NodeType) (*NodeType, error)
	// BatchCreateNodeTypes 批量创建NodeType
	BatchCreateNodeTypes(context.Context, []*NodeType) ([]*NodeType, error)
	// GetNodeTypeById 根据节点类型ID获取节点类型
	GetNodeTypeById(context.Context, int) (*NodeType, error)
	// GetNodeTypesByIds 根据节点类型ID列表获取节点类型列表
	GetNodeTypesByIds(context.Context, []int) ([]*NodeType, error)
	// UpdateNodeTypeById 根据节点类型ID更新节点类型
	UpdateNodeTypeById(context.Context, int, *NodeType) (bool, error)
	// DeleteNodeTypeById 根据节点类型ID删除节点类型，仅当删除成功时返回True
	DeleteNodeTypeById(ctx context.Context, id int) (bool, error)
}

type NodeTypeUsecase struct {
	tm     tx.Transaction
	repo   NodeTypeRepo
	logger *log.Helper
}

func NewNodeTypeUsecase(tm tx.Transaction, repo NodeTypeRepo, logger log.Logger) *NodeTypeUsecase {
	return &NodeTypeUsecase{tm: tm, repo: repo, logger: log.NewHelper(logger)}
}

// CreateNodeType 创建节点类型
func (uc *NodeTypeUsecase) CreateNodeType(ctx context.Context, nodeType *NodeType) (*NodeType, error) {
	return uc.repo.CreateNodeType(ctx, nodeType)
}

// GetNodeTypeById 根据ID查询节点类型
func (uc *NodeTypeUsecase) GetNodeTypeById(ctx context.Context, id int) (*NodeType, error) {
	return uc.repo.GetNodeTypeById(ctx, id)
}

// UpdateNodeTypeById 根据ID更新节点类型
func (uc *NodeTypeUsecase) UpdateNodeTypeById(ctx context.Context, id int, nodeType *NodeType) error {
	_, err := uc.repo.UpdateNodeTypeById(ctx, id, nodeType)
	return err
}

// DeleteNodeTypeById 根据ID删除节点类型
func (uc *NodeTypeUsecase) DeleteNodeTypeById(ctx context.Context, id int) error {
	// FIXME 删除节点类型之前，需要确保已经没有地图在使用该节点类型
	nodeType, err := uc.repo.GetNodeTypeById(ctx, id)
	if err != nil {
		return err
	}
	if nodeType == nil {
		return data_errors.NewResourceNotFoundError("node_type id: %d not found", id)
	}
	if data.IsDefault(nodeType.IsDefault) {
		return data_errors.NewDefaultResourceForbiddenError("delete a default NodeType is not allowed")
	}
	_, err = uc.repo.DeleteNodeTypeById(ctx, id)
	return err
}
