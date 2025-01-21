package data

import (
	"context"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
	"github.com/go-kratos/kratos/v2/log"
)

type sysRoleRepo struct {
	data *Data
	log  *log.Helper
}

func NewSysRoleRepo(data *Data, logger log.Logger) biz.SysRoleRepo {
	return &sysRoleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (s *sysRoleRepo) ListPage(ctx context.Context, name, key string, status, num, size int32) ([]*model.SysRole, error) {
	q := s.data.Query(ctx).SysRole
	db := q.WithContext(ctx)
	if name != "" {
		db = db.Where(q.RoleName.Like(buildLikeValue(name)))
	}
	if key != "" {
		db = db.Where(q.RoleKey.Eq(key))
	}
	if status != 0 {
		db = db.Where(q.Status.Eq(status))
	}
	limit, offset := convertPageSize(num, size)
	return db.Order(q.RoleSort.Desc()).Limit(limit).Offset(offset).Find()
}

func (s *sysRoleRepo) Count(ctx context.Context, name, key string, status int32) (int32, error) {
	db := s.data.db.SysRole.WithContext(ctx)
	if name != "" {
		db = db.Where(s.data.db.SysRole.RoleName.Like(name))
	}
	if key != "" {
		db = db.Where(s.data.db.SysRole.RoleKey.Like(key))
	}
	if status != 0 {
		db = db.Where(s.data.db.SysRole.Status.Eq(status))
	}
	count, err := db.Count()
	return int32(count), err
}

func (s *sysRoleRepo) FindByID(ctx context.Context, id int64) (*model.SysRole, error) {
	db := s.data.db.SysRole.WithContext(ctx)

	return db.Where(s.data.db.SysRole.ID.Eq(id)).First()
}
