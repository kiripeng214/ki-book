package server

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/middleware"
	"github.com/go-kratos/kratos/v2/transport"
	"google.golang.org/grpc/metadata"
)

type authReq struct {
	Secret string `json:"secret"`
}

func NewAuthServer(handler middleware.Handler) middleware.Handler {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		if md, ok := metadata.FromIncomingContext(ctx); ok {
			fmt.Println("dd", md)
		} else if header, ok := transport.FromServerContext(ctx); ok {
			fmt.Println("ff", header)
		}
		reply, err := handler(ctx, req)
		return reply, err
	}
}
