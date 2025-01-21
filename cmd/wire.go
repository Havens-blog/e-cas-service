//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package cmd

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/Havens-blog/e-cas-service/configs/conf"
	"github.com/Havens-blog/e-cas-service/internal/biz"
	"github.com/Havens-blog/e-cas-service/internal/data"
	"github.com/Havens-blog/e-cas-service/internal/server"
	"github.com/Havens-blog/e-cas-service/internal/service"
)

// wireApp init kratos application.
func wireApp(*conf.Server, *conf.Data, *conf.Consul, *conf.Global, *conf.Session, log.Logger) (*kratos.App, error) {
	panic(wire.Build(server.ProviderSet, data.ProviderSet, biz.ProviderSet, service.ProviderSet, newApp))
}
