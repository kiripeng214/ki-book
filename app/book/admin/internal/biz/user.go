package biz

import (
	"context"
	"fmt"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	usV1 "ki-book/api/user/service/v1"
)

var _ UserRepo = (*UserUseCase)(nil)

type User struct {
	Name     string `json:"name"`
	Password string `json:"password"`
	Phone    string `json:"phone"`
	Result   bool   `json:"result"`
	Id       int64  `json:"id"`
	Age      uint64 `json:"age"`
}

type UserRepo interface {
	Login(ctx context.Context, user *User) (*User, error)
}

type UserUseCase struct {
	repo UserRepo
	us   usV1.UserClient
	log  *log.Helper
}

func NewUserCase(repo UserRepo, logger log.Logger, client usV1.UserClient) *UserUseCase {
	logs := log.NewHelper(log.With(logger, "module", "usecase/interface"))
	return &UserUseCase{
		repo: repo,
		us:   client,
		log:  logs,
	}
}

func (u *UserUseCase) Login(ctx context.Context, user *User) (*User, error) {
	reply, err := u.us.VerifyPassword(ctx, &usV1.VerifyPasswordReq{
		Name:     user.Name,
		Password: user.Password,
	})
	if err != nil {

		fmt.Println(errors.Code(err))
		return nil, err
	}

	getUser, err := u.us.GetUser(ctx, &usV1.GetUserReq{Id: reply.GetId()})
	if err != nil {
		return nil, err
	}

	return &User{
		Result: reply.GetResult(),
		Id:     reply.GetId(),
		Name:   getUser.GetName(),
		Phone:  getUser.GetPhone(),
		Age:    getUser.GetAge(),
	}, nil
}
