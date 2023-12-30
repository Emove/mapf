package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

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
