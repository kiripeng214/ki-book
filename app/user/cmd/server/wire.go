//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	biz2 "ki-book/app/user/internal/biz"
	conf2 "ki-book/app/user/internal/conf"
	data2 "ki-book/app/user/internal/data"
	server2 "ki-book/app/user/internal/server"
	service2 "ki-book/app/user/internal/service"
)

// initApp init kratos application.
func initApp(*conf2.Server, *conf2.Registry, *conf2.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server2.ProviderSet, data2.ProviderSet, biz2.ProviderSet, service2.ProviderSet, newApp))
}
