package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/app/warehouse/internal/biz"
)

type affixNodeRepo struct {
	data   *Data
	logger *log.Helper
}

func NewAffixNodeRepo(data *Data, logger log.Logger) (biz.AffixNodeRepo, error) {
	return &affixNodeRepo{
		data:   data,
		logger: log.NewHelper(logger),
	}, nil
}
