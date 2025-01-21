// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package cmd

import (
	"github.com/Havens-blog/e-cas-service/configs/conf"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data"
	"github.com/Havens-blog/e-cas-service/internal/server"
	"github.com/Havens-blog/e-cas-service/internal/service"
	"github.com/Havens-blog/e-cas-service/pkg/consul"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
)

import (
	_ "github.com/Havens-blog/e-cas-service/internal/biz"
	_ "go.uber.org/automaxprocs"
)

// Injectors from wire.go:

// wireApp init kratos application.
func wireApp(confServer *conf.Server, confData *conf.Data, confConsul *conf.Consul, global *conf.Global, session *conf.Session, logger log.Logger) (*kratos.App, error) {
	dataData, err := data.NewData(confData, logger)
	if err != nil {
		return nil, err
	}
	repo := data.NewDataRepo(dataData, logger)
	httpUsecase := biz.NewHttpUsecase(repo, logger)
	httpService := service.NewHttpService(httpUsecase, logger)
	sysRoleRepo := data.NewSysRoleRepo(dataData, logger)
	sysRoleUseCase := biz.NewSysRoleUseCase(sysRoleRepo, logger)
	rolesService := service.NewRolesService(logger, sysRoleUseCase)
	sysMenuRepo := data.NewSysMenuRepo(dataData, logger)
	sysMenuUseCase := biz.NewSysMenuUseCase(sysMenuRepo, logger)
	sysRoleMenuRepo := data.NewSysRoleMenuRepo(dataData, logger)
	sysRoleMenuUseCase := biz.NewSysRoleMenuUseCase(sysRoleMenuRepo, logger)
	menusService := service.NewMenusService(sysMenuUseCase, sysRoleMenuUseCase, logger)
	sysUserRepo := data.NewSysUserRepo(dataData, logger)
	sysUserUseCase := biz.NewSysUserUseCase(sysUserRepo, confServer, logger)
	sysuserService := service.NewSysuserService(sysUserUseCase, sysRoleUseCase, sysRoleMenuUseCase, logger)
	httpServer := server.NewHTTPServer(confServer, session, httpService, rolesService, menusService, sysuserService, logger)
	grpcUsecase := biz.NewGrpcUsecase(repo, logger)
	grpcService := service.NewGrpcService(grpcUsecase, logger)
	grpcServer := server.NewGRPCServer(confServer, grpcService, logger)
	registrar, err := consul.NewRegistry(confConsul)
	if err != nil {
		return nil, err
	}
	app := newApp(logger, httpServer, grpcServer, registrar, global)
	return app, nil
}
