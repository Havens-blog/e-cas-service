package service

import (
	"context"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/go-kratos/kratos/v2/log"

	v1 "github.com/Havens-blog/e-cas-service/api/grpc/v1"
)

type GrpcService struct {
	v1.UnimplementedGrpcServer
	log *log.Helper
	uc  *biz.GrpcUsecase
}

func NewGrpcService(uc *biz.GrpcUsecase, lg log.Logger) *GrpcService {
	return &GrpcService{
		uc:  uc,
		log: log.NewHelper(lg),
	}
}
func (s *GrpcService) GetUserInfo(ctx context.Context, req *v1.UserInfoRequest) (*v1.UserInfoResponse, error) {
	return &v1.UserInfoResponse{
		Code: 21312,
	}, nil
}
