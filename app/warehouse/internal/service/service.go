package service

import (
	"github.com/google/wire"
	v1 "mapf/api/warehouse/v1"
	"mapf/app/warehouse/internal/biz"
	"mapf/internal/data"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWarehouseService)

type WarehouseService struct {
	v1.UnimplementedWarehouseServiceServer

	wuc   *biz.WarehouseUsecase
	ntuc  *biz.NodeTypeUsecase
	nuc   *biz.NodeUsecase
	nciuc *biz.NodeConfigItemUsecase
	anuc  *biz.AffixNodeUsecase
}

// NewWarehouseService new a warehouse service.
func NewWarehouseService(
	wuc *biz.WarehouseUsecase,
	ntuc *biz.NodeTypeUsecase,
	nuc *biz.NodeUsecase,
	nciuc *biz.NodeConfigItemUsecase,
	anuc *biz.AffixNodeUsecase,
) *WarehouseService {
	return &WarehouseService{
		wuc:   wuc,
		ntuc:  ntuc,
		nuc:   nuc,
		nciuc: nciuc,
		anuc:  anuc,
	}
}

func ConvertModelStatus2protobufEnableStatusEnum(status int) v1.EnableStatus {
	if status == data.DisableStatus {
		return v1.EnableStatus_DISABLE
	}
	return v1.EnableStatus_ENABLE
}

func ConvertModelIsDefault2protoIsDefaultEnum(isDefault int) v1.DefaultStatus {
	if isDefault == data.IsDefaultFalse {
		return v1.DefaultStatus_NOT_DEFAULT
	}
	return v1.DefaultStatus_IS_DEFAULT
}

func ConvertProtobufEnableStatus2modelStatus(status v1.EnableStatus) int {
	if status == v1.EnableStatus_ENABLE {
		return data.EnableStatus
	}
	return data.DisableStatus
}
