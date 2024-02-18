package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
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
