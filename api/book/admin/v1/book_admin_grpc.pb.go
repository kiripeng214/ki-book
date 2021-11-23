// Code generated by protoc-gen-go-grpc. DO NOT EDIT.

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BookAdminClient is the client API for BookAdmin service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BookAdminClient interface {
	Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error)
}

type bookAdminClient struct {
	cc grpc.ClientConnInterface
}

func NewBookAdminClient(cc grpc.ClientConnInterface) BookAdminClient {
	return &bookAdminClient{cc}
}

func (c *bookAdminClient) Login(ctx context.Context, in *LoginReq, opts ...grpc.CallOption) (*LoginReply, error) {
	out := new(LoginReply)
	err := c.cc.Invoke(ctx, "/book.admin.v1.BookAdmin/Login", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BookAdminServer is the server API for BookAdmin service.
// All implementations must embed UnimplementedBookAdminServer
// for forward compatibility
type BookAdminServer interface {
	Login(context.Context, *LoginReq) (*LoginReply, error)
	mustEmbedUnimplementedBookAdminServer()
}

// UnimplementedBookAdminServer must be embedded to have forward compatible implementations.
type UnimplementedBookAdminServer struct {
}

func (UnimplementedBookAdminServer) Login(context.Context, *LoginReq) (*LoginReply, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Login not implemented")
}
func (UnimplementedBookAdminServer) mustEmbedUnimplementedBookAdminServer() {}

// UnsafeBookAdminServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BookAdminServer will
// result in compilation errors.
type UnsafeBookAdminServer interface {
	mustEmbedUnimplementedBookAdminServer()
}

func RegisterBookAdminServer(s grpc.ServiceRegistrar, srv BookAdminServer) {
	s.RegisterService(&BookAdmin_ServiceDesc, srv)
}

func _BookAdmin_Login_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(LoginReq)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BookAdminServer).Login(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/book.admin.v1.BookAdmin/Login",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BookAdminServer).Login(ctx, req.(*LoginReq))
	}
	return interceptor(ctx, in, info, handler)
}

// BookAdmin_ServiceDesc is the grpc.ServiceDesc for BookAdmin service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BookAdmin_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "book.admin.v1.BookAdmin",
	HandlerType: (*BookAdminServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Login",
			Handler:    _BookAdmin_Login_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "api/book/admin/v1/book_admin.proto",
}
