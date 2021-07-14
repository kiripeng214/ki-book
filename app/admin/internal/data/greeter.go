package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	biz2 "ki-book/app/admin/internal/biz"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz2.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) CreateGreeter(ctx context.Context, g *biz2.Greeter) error {
	return nil
}

func (r *greeterRepo) UpdateGreeter(ctx context.Context, g *biz2.Greeter) error {
	return nil
}
