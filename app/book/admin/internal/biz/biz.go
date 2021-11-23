package biz

import (
	"context"
	etcdRegistry "github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/google/wire"
	etcd "go.etcd.io/etcd/client/v3"
	usV1 "ki-book/api/user/service/v1"
	"ki-book/app/book/admin/internal/conf"
)

var ProviderSet3 = wire.NewSet(
	NewUserCase,
	NewDiscovery,
	NewUserServiceClient,
)

func NewUserServiceClient(r registry.Discovery) usV1.UserClient {
	clientConn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///ki.user.service"),
		grpc.WithDiscovery(r),
	)
	if err != nil {
		panic(err)
	}
	return usV1.NewUserClient(clientConn)
}

func NewDiscovery(conf *conf.Registry) registry.Discovery {
	client, err := etcd.New(etcd.Config{
		Endpoints: conf.Etcd.GetAddress(),
	})
	if err != nil {
		panic(err)
	}
	return etcdRegistry.New(client)
}
