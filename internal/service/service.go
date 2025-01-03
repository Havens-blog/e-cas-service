package service

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	v1 "github.com/Havens-blog/e-cas-service/api/grpc/v1"
	"github.com/Havens-blog/e-cas-service/internal/biz"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHttpService, NewGrpcService)

type HttpService struct {
	log *log.Helper
	uc  *biz.HttpUsecase
}

type GrpcService struct {
	v1.UnimplementedGrpcServer
	log *log.Helper
	uc  *biz.GrpcUsecase
}

func NewHttpService(uc *biz.HttpUsecase, lg log.Logger) *HttpService {
	return &HttpService{
		uc:  uc,
		log: log.NewHelper(lg),
	}
}

func NewGrpcService(uc *biz.GrpcUsecase, lg log.Logger) *GrpcService {
	return &GrpcService{
		uc:  uc,
		log: log.NewHelper(lg),
	}
}
