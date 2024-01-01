package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	warehouse "mapf/api/warehouse/v1"
	"mapf/app/web/internal/data"
)

func NewWarehouseUsecase(data *data.Data, logger log.Logger) *WarehouseUsecase {
	return &WarehouseUsecase{whc: data.Whc, logger: log.NewHelper(logger)}
}

type WarehouseUsecase struct {
	data   *data.Data
	whc    warehouse.WarehouseServiceClient
	logger *log.Helper
}
