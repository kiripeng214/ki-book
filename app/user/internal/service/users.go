package service

import (
    "context"

    pb "ki-book/api/user/service/v1"
)



func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	return &pb.GetUserReply{}, nil
}
func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	return &pb.CreateUserReply{}, nil
}
func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
	return &pb.VerifyPasswordReply{}, nil
}
