package biz

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	pbAny "github.com/Havens-blog/e-cas-service/api/any/v1"
	"github.com/Havens-blog/e-cas-service/internal/data/model"
)

// ProviderSet is biz providers.
var ProviderSet = wire.NewSet(NewHttpUsecase, NewGrpcUsecase, NewSysRoleUseCase, NewSysMenuUseCase,
	NewSysRoleMenuUseCase,
	NewSysUserUseCase,
)

type GrpcUsecase struct {
	repo Repo
	log  *log.Helper
}

func NewGrpcUsecase(repo Repo, lg log.Logger) *GrpcUsecase {
	return &GrpcUsecase{
		repo: repo,
		log:  log.NewHelper(lg),
	}
}

type Repo interface {
	MysqlInterface
	ESInterface
}

type MysqlInterface interface {
	ListAllUser(context.Context, *model.User, *pbAny.PageRequest) ([]*model.User, int64, error)
	FirstUser(context.Context, *model.SysUser) (*model.SysUser, error)
	UpdateUserInfo(context.Context, *model.User) error
	UpdatePassword(context.Context, *model.User) error

	UpdateSetting(context.Context, *model.Setting) error
	GetSetting(context.Context, *model.Setting) (*model.Setting, error)

	AddRoleMenu(context.Context, *model.Menu) error
	DeleteRoleMenuByID(context.Context, *model.Menu) error
	UpdateRoleMenu(context.Context, *model.Menu) error
	ListAllMenu(context.Context) ([]*model.Menu, error)
	ListRoleMenu(context.Context, *model.Menu) ([]*model.Menu, error)

	ListRoleCasbin(context.Context, *model.CasbinRule, *pbAny.PageRequest) ([]*model.CasbinRule, int64, error)
	AddRoleCasbin(context.Context, *model.CasbinRule) error
	UpdateRoleCasbin(context.Context, *model.CasbinRule) error
	DeleteRoleCasbin(context.Context, *model.CasbinRule) error

	FirstDebugPerf(context.Context, *model.DebugPerf) (*model.DebugPerf, error)
	AddDebugPerf(context.Context, *model.DebugPerf) error
	UpdateDebugPerf(context.Context, *model.DebugPerf) error
	DeleteDebugPerf(context.Context, *model.DebugPerf) error
	TransactionDebugPerf(context.Context, *model.DebugPerf) error
}

type ESInterface interface {
	ESInsertDoc(context.Context, string, interface{}) error
}
