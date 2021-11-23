//go:build wireinject
// +build wireinject

// The build tag makes sure the stub is not built in the final build.

package main

import (
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	biz2 "ki-book/app/book/admin/internal/biz"
	conf2 "ki-book/app/book/admin/internal/conf"
	data2 "ki-book/app/book/admin/internal/data"
	server2 "ki-book/app/book/admin/internal/server"
	service2 "ki-book/app/book/admin/internal/service"
)

// initApp init kratos application.
func initApp(*conf2.Server, *conf2.Registry, *conf2.Data, log.Logger) (*kratos.App, func(), error) {
	panic(wire.Build(server2.ProviderSet1, data2.ProviderSet2, biz2.ProviderSet3, service2.ProviderSet, newApp))
}
