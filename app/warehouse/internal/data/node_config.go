package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
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
