package service

import (
	"github.com/google/wire"
	robot "mapf/api/robot/v1"
	"mapf/app/robot/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewRobotService)

// RobotService is a robot service.
type RobotService struct {
	robot.UnimplementedRobotServiceServer

	ruc *biz.RobotUsecase
}

// NewRobotService new a robot service.
func NewRobotService(ruc *biz.RobotUsecase) *RobotService {
	return &RobotService{ruc: ruc}
}
