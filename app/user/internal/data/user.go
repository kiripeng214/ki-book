package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"ki-book/app/user/internal/biz"
	"ki-book/app/user/internal/pkg/util"
)

var _ biz.UserRepo = (*userRepo)(nil)

type userRepo struct {
	data *Data
	log  *log.Helper
}

func NewUserRepo(data *Data, logger log.Logger) biz.UserRepo {
	return &userRepo{
		data: data,
		log:  log.NewHelper(log.With(logger, "module", "data/server-service")),
	}
}

func (u2 userRepo) CreateUser(ctx context.Context, u *biz.User) (*biz.User, error) {
	hs, err := util.HashPassword(u.Secret)
	if err != nil {
		return nil, err
	}
	save, err := u2.data.db.User.
		Create().
		SetName(u.Name).
		SetSecret(hs).
		SetAge(u.Age).
		SetPhone(u.Phone).Save(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.User{
		Id:    save.ID,
		Name:  save.Name,
		Age:   save.Age,
		Phone: save.Phone,
	}, nil

}

func (u2 userRepo) GetUser(ctx context.Context, id int64) (*biz.User, error) {
	user, err := u2.data.db.User.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	return &biz.User{Id: user.ID, Name: user.Name, Age: user.Age, Phone: user.Phone}, err
}

func (u2 userRepo) VerifyPassword(ctx context.Context, u *biz.User) (bool, error) {
	panic("implement me")
}
