package biz

import (
	"context"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/go-kratos/kratos/v2/log"
)

type SysRoleRepo interface {
	ListPage(ctx context.Context, name, key string, status, num, size int32) ([]*model.SysRole, error)
	Count(ctx context.Context, name, key string, status int32) (int32, error)
	FindByID(ctx context.Context, id int64) (*model.SysRole, error)
}

type SysRoleUseCase struct {
	repo SysRoleRepo
	log  *log.Helper
}

func (s *SysRoleUseCase) ListPage(ctx context.Context, roleName, roleKey string, status, page, size int32) ([]*model.SysRole, int32, error) {
	total, err := s.repo.Count(ctx, roleName, roleKey, status)
	if err != nil {
		return nil, 0, err
	}
	roleList, err := s.repo.ListPage(ctx, roleName, roleKey, status, page, size)
	return roleList, total, err
}

func (s *SysRoleUseCase) GetRole(ctx context.Context, id int64) (*model.SysRole, error) {
	return s.repo.FindByID(ctx, id)
}

func NewSysRoleUseCase(repo SysRoleRepo, logger log.Logger) *SysRoleUseCase {
	return &SysRoleUseCase{repo: repo, log: log.NewHelper(logger)}
}
