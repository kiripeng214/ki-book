package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"ki-book/app/book/admin/internal/biz"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func (u userRepo) Login(ctx context.Context, user *biz.User) (*biz.User, error) {
	panic("implement me")
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/user")),
	}
}
