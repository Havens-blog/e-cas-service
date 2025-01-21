package service

import (
	"github.com/google/wire"
)

// ProviderSet is service providers.
var ProviderSet = wire.NewSet(NewHttpService, NewGrpcService, NewRolesService, NewMenusService,
	NewSysuserService)
