package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

type nodeRepo struct {
	data   *Data
	logger *log.Helper
}

func NewNodeRepo(data *Data, logger log.Logger) (biz.NodeRepo, error) {
	return &nodeRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
