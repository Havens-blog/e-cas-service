package biz

import (
	pb "github.com/Havens-blog/e-cas-service/api/http/cas/v1"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type SysRoleMenuRepo interface {
	FindMenuByRoleId(ctx *gin.Context, id int64) ([]*model.SysMenu, error)
	GetPermission(ctx *gin.Context, id int64) ([]string, error)
	SelectMenuRole(c *gin.Context, name string) ([]*pb.MenuTree, error)
}

type SysRoleMenuUseCase struct {
	repo SysRoleMenuRepo
	log  *log.Helper
}

func NewSysRoleMenuUseCase(repo SysRoleMenuRepo, logger log.Logger) *SysRoleMenuUseCase {
	return &SysRoleMenuUseCase{repo: repo, log: log.NewHelper(logger)}
}

func (s *SysRoleMenuUseCase) GetPermission(ctx *gin.Context, roleID int64) ([]string, error) {
	return s.repo.GetPermission(ctx, roleID)
}

func (s *SysRoleMenuUseCase) FindMenuByRoleId(ctx *gin.Context, roleID int64) ([]*model.SysMenu, error) {
	return s.repo.FindMenuByRoleId(ctx, roleID)
}

func (s *SysRoleMenuUseCase) SelectMenuRole(c *gin.Context, name string) ([]*pb.MenuTree, error) {
	return s.repo.SelectMenuRole(c, name)

}
