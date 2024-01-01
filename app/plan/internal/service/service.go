package service

import (
	"github.com/google/wire"
	plan "mapf/api/plan/v1"
	"mapf/app/plan/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewPlanServiceServer)

// PlanService is a robot service.
type PlanService struct {
	plan.UnimplementedPlanServiceServer
}

// NewPlanServiceServer new a plan service.
func NewPlanServiceServer(uc *biz.GreeterUsecase) *PlanService {
	return &PlanService{}
}
