package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

type warehouseRepo struct {
	data   *Data
	logger *log.Helper
}

func NewWarehouseRepo(data *Data, logger log.Logger) (biz.WarehouseRepo, error) {
	return &warehouseRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
