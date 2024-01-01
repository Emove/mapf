package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	v1 "mapf/api/web/v1"
	"mapf/app/web/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewWebService)

type WebService struct {
	v1.UnimplementedWebServiceServer

	log *log.Helper
	whc *biz.WarehouseUsecase
}

func NewWebService(
	whc *biz.WarehouseUsecase,
	logger log.Logger) *WebService {
	return &WebService{
		log: log.NewHelper(log.With(logger, "module", "service/web")),
		whc: whc,
	}
}
