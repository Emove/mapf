package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

type nodeConfigItemRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeConfigItemRepo(data *Data, logger log.Logger) (biz.NodeConfigItemRepo, error) {
	return &nodeConfigItemRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
