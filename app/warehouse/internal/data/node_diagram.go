package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
)

type nodeDiagramRepo struct {
	*data.Pager
	data   *Data
	logger *log.Helper
}

func NewNodeDiagramRepo(data *Data, logger log.Logger) (biz.NodeDiagramRepo, error) {
	return &nodeDiagramRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
