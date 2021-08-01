package server

import (
	etcdRegistry "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcd "go.etcd.io/etcd/client/v3"
	"ki-book/app/user/internal/conf"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(NewGRPCServer, NewRegistrar)

func NewRegistrar(conf *conf.Registry) registry.Registrar {
	client, err := etcd.New(etcd.Config{
		Endpoints: conf.Etcd.GetAddress(),
	})
	if err != nil {
		panic(err)
	}
	return etcdRegistry.New(client)
}
