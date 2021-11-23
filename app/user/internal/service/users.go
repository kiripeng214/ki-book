package service

import (
	"context"
	pb "ki-book/api/user/service/v1"
	"ki-book/app/user/internal/biz"
)

func (s *UserService) GetUser(ctx context.Context, req *pb.GetUserReq) (*pb.GetUserReply, error) {
	user, err := s.uc.GetUser(ctx, req.GetId())
	if err != nil {
		return nil, err
	}
	return &pb.GetUserReply{
		Id:    user.Id,
		Name:  user.Name,
		Age:   user.Age,
		Phone: user.Phone,
	}, nil
}

func (s *UserService) CreateUser(ctx context.Context, req *pb.CreateUserReq) (*pb.CreateUserReply, error) {
	user, err := s.uc.CreateUser(ctx, &biz.User{
		Name:   req.GetName(),
		Secret: req.GetPassword(),
		Age:    req.GetAge(),
		Phone:  req.GetPhone(),
	})
	if err != nil {
		return nil, err
	}
	return &pb.CreateUserReply{
		Id: user.Id,
	}, err
}

func (s *UserService) VerifyPassword(ctx context.Context, req *pb.VerifyPasswordReq) (*pb.VerifyPasswordReply, error) {
	info, err := s.uc.VerifyPassword(ctx, &biz.User{
		Name:   req.GetName(),
		Secret: req.GetPassword(),
	})
	return &pb.VerifyPasswordReply{
		Result: info.IsExist,
		Id:     info.Id,
	}, err
}
