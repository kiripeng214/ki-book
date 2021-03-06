package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	conf2 "ki-book/app/interface/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO warpped database client
}

// NewData .
func NewData(c *conf2.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		logger.Log(log.LevelInfo, "closing the data resources")
	}
	return &Data{}, cleanup, nil
}
