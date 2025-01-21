package biz

import (
	"github.com/Havens-blog/e-cas-service/configs/conf"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type SysUserRepo interface {
	FindByUsername(ctx *gin.Context, name string) (*model.SysUser, error)
}

type SysUserUseCase struct {
	userRepo    SysUserRepo
	severConfig *conf.Server
	log         *log.Helper
}

func (s *SysUserUseCase) FindUserByUsername(c *gin.Context, username string) (*model.SysUser, error) {
	return s.userRepo.FindByUsername(c, username)
}

func NewSysUserUseCase(userRepo SysUserRepo, severConfig *conf.Server, logger log.Logger) *SysUserUseCase {
	return &SysUserUseCase{
		userRepo:    userRepo,
		severConfig: severConfig,
		log:         log.NewHelper(logger),
	}
}
