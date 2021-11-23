package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"ki-book/app/book/admin/internal/conf"
)

var ProviderSet2 = wire.NewSet(NewData, NewUserRepo)

type Data struct {
	logs *log.Helper
}

func NewData(conf *conf.Data, logger log.Logger) (*Data, error) {
	l := log.NewHelper(log.With(logger, "moudle", "ki-admin/data"))
	return &Data{logs: l}, nil
}
