package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"ki-book/app/user/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log *log.Helper
}

func NewUserRepo(data *Data,logger log.Logger) biz.UserRepo{
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (u2 userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	panic("implement me")
}

func (u2 userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	panic("implement me")
}

func (u2 userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	panic("implement me")
}

