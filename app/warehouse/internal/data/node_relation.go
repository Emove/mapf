package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

type nodeRelationRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeRelationRepo(data *Data, logger log.Logger) (biz.NodeRelationRepo, error) {
	return &nodeRelationRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
