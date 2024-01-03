package service

import (
	"github.com/google/wire"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWarehouseService)

type WarehouseService struct {
	v1.UnimplementedWarehouseServiceServer

	uc *biz.WarehouseUsecase
}

// NewWarehouseService new a warehouse service.
func NewWarehouseService(uc *biz.WarehouseUsecase) *WarehouseService {
	return &WarehouseService{uc: uc}
}
