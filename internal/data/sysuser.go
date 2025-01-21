package data

import (
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/gin-gonic/gin"
	"github.com/go-kratos/kratos/v2/log"
)

type sysUserRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysUserRepo(data *Data, logger log.Logger) biz.SysUserRepo {
	return &sysUserRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *sysUserRepo) FindByUsername(ctx *gin.Context, name string) (*model.SysUser, error) {
	query := s.data.Query(ctx).SysUser
	return query.WithContext(ctx).Where(query.Username.Eq(name)).First()
}
