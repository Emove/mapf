package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

type nodeConfigRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeConfigRepo(data *Data, logger log.Logger) (biz.NodeConfigRepo, error) {
	return &nodeConfigRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
