package biz

import (
	"github.com/go-kratos/kratos/v2/log"
	"mapf/internal/data"
)

type Warehouse struct {
	*data.Model
	Name      string `gorm:"type:varchar(64)"`
	IsDefault int    `gorm:"type:smallint;default:0"`
	Status    int    `gorm:"type:smallint;default:1"`
}

type WarehouseRepo interface {
}

type WarehouseUsecase struct {
	repo   WarehouseRepo
	logger *log.Helper
}

func NewWarehouseUsecase(repo WarehouseRepo, logger log.Logger) *WarehouseUsecase {
	return &WarehouseUsecase{repo: repo, logger: log.NewHelper(logger)}
}