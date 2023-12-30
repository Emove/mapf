package service

import (
	"context"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
)

type WarehouseService struct {
	v1.UnimplementedWarehouseServer

	uc *biz.WarehouseUsecase
}

// NewWarehouseService new a warehouse service.
func NewWarehouseService(uc *biz.WarehouseUsecase) *WarehouseService {
	return &WarehouseService{uc: uc}
}

// CreateWarehouse implements warehouse.CreateWarehouse.
func (s *GreeterService) CreateWarehouse(ctx context.Context, in *v1.CreateWarehouseRequest) (*v1.CreateWarehouseResponse, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: in.Name})
	if err != nil {
		return nil, err
	}
	return &v1.CreateWarehouseResponse{Message: "Hello " + g.Hello}, nil
}
