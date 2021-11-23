package service

import (
	"context"
	"ki-book/app/book/admin/internal/biz"

	pb "ki-book/api/book/admin/v1"
)

type BookAdminService struct {
	pb.UnimplementedBookAdminServer
	uc *biz.UserUseCase
}

func NewBookAdminService(useCase *biz.UserUseCase) *BookAdminService {
	return &BookAdminService{
		uc: useCase,
	}
}

func (s *BookAdminService) Login(ctx context.Context, req *pb.LoginReq) (*pb.LoginReply, error) {

	login, err := s.uc.Login(ctx, &biz.User{
		Name:     req.Username,
		Password: req.Password,
	})
	if err != nil {
		return nil, err
	}

	return &pb.LoginReply{
		Id:    login.Id,
		Phone: login.Phone,
		Name:  login.Name,
		Age:   login.Age,
	}, nil
}
