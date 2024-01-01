package service

import (
	"context"
	"github.com/google/wire"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWarehouseService, NewGreeterService)

type WarehouseService struct {
	v1.UnimplementedWarehouseServiceServer

	uc *biz.WarehouseUsecase
}

// NewWarehouseService new a warehouse service.
func NewWarehouseService(uc *biz.WarehouseUsecase) *WarehouseService {
	return &WarehouseService{uc: uc}
}

// CreateWarehouse implements warehouse.CreateWarehouse.
func (s *GreeterService) CreateWarehouse(ctx context.Context, in *v1.CreateWarehouseRequest) (*v1.CreateWarehouseResponse, error) {
	//g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	//if err != nil {
	//	return nil, err
	//}
	return &v1.CreateWarehouseResponse{}, nil
}
