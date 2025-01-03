package server

import (
	"github.com/google/wire"

	"github.com/Havens-blog/e-cas-service/pkg/consul"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewHTTPServer, consul.NewRegistry)
