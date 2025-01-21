package biz

import (
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type SysMenuRepo interface {
	FindById(ctx *gin.Context, id int64) (*model.SysMenu, error)
	DeleteRoleMenuByID(ctx *gin.Context, id int64) error
	AddRoleMenu(ctx *gin.Context, req *model.SysMenu) error
}

type SysMenuUseCase struct {
	repo SysMenuRepo
	log  *log.Helper
}

func NewSysMenuUseCase(repo SysMenuRepo, logger log.Logger) *SysMenuUseCase {
	return &SysMenuUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s *SysMenuUseCase) GetMenu(ctx *gin.Context, id int64) (*model.SysMenu, error) {
	return s.repo.FindById(ctx, id)
}

func (s *SysMenuUseCase) DeleteRoleMenuByID(ctx *gin.Context, id int64) error {
	return s.repo.DeleteRoleMenuByID(ctx, id)
}

func (s *SysMenuUseCase) AddRoleMenu(ctx *gin.Context, req *model.SysMenu) error {
	return s.repo.AddRoleMenu(ctx, req)
}
