package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
)

type nodeTagRepo struct {
	*data.Pager
	data   *Data
	logger *log.Helper
}

func NewNodeTagRepo(data *Data, logger log.Logger) (biz.NodeTagRepo, error) {
	return &nodeTagRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
