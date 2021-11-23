package server

import (
	etcdRegistry "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/google/wire"
	etcd "go.etcd.io/etcd/client/v3"
	"ki-book/app/book/admin/internal/conf"
	"time"
)

var ProviderSet1 = wire.NewSet(NewGRPCServer, NewRegistrar, NewHTTPServer)

func NewRegistrar(conf *conf.Registry) registry.Registrar {

	client, err := etcd.New(etcd.Config{
		Endpoints:         conf.Etcd.GetAddress(),
		DialTimeout:       5 * time.Second,
		DialKeepAliveTime: 3 * time.Second,
	})

	if err != nil {
		panic(err)
	}
	return etcdRegistry.New(client)
}
