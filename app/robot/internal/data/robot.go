package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/robot/internal/biz"
)

type robotRepo struct {
	logger *log.Helper
}

func NewRobotRepo(data *Data, logger log.Logger) biz.RobotRepo {
	return &robotRepo{
		logger: log.NewHelper(logger),
	}
}
