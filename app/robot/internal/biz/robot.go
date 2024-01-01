package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"mapf/internal/data"
)

// Robot is a Robot model.
type Robot struct {
	*data.Model
}

// RobotRepo is a Robot repo.
type RobotRepo interface {
}

// RobotUsecase is a Robot usecase.
type RobotUsecase struct {
	repo RobotRepo
	log  *log.Helper
}

// NewRobotUsecase new a Robot usecase.
func NewRobotUseCase(repo RobotRepo, logger log.Logger) *RobotUsecase {
	return &RobotUsecase{repo: repo, log: log.NewHelper(logger)}
}

// CreateRobot creates a Robot, and returns the new Robot.
func (uc *RobotUsecase) CreateRobot(ctx context.Context, r *Robot) (*Greeter, error) {
	return nil, nil
}
