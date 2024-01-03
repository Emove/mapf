package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
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

func (repo *nodeTypeRepo) BatchCreateNodeTypes(ctx context.Context, nodeTypes []*biz.NodeType) ([]*biz.NodeType, error) {
	db := repo.data.DB(ctx)
	for i := 0; i < len(nodeTypes); i++ {
		if err := db.Create(nodeTypes[i]).Error; err != nil {
			return nodeTypes, err
		}
	}
	return nodeTypes, nil
}
