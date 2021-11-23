package data

import (
	"github.com/go-kratos/kratos/v2/log"
	_ "github.com/go-sql-driver/mysql"
	"github.com/google/wire"
	"ki-book/app/user/internal/conf"
	"ki-book/app/user/internal/data/ent"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewEntClient, NewUserRepo)

// Data .
type Data struct {
	db *ent.Client
}

func NewEntClient(c *conf.Data, logger log.Logger) *ent.Client {
	logs := log.NewHelper(log.With(logger, "moudle", "user-service/data/ent"))
	client, err := ent.Open(
		c.Database.Driver,
		c.Database.Source,
	)
	if err != nil {
		logs.Fatalf("failed opening connection to db: %v", err)
	}
	//if err := client.Schema.Create(context.Background(), migrate.WithForeignKeys(false)); err != nil {
	//	logs.Fatalf("failed creating schema resources: %v", err)
	//}
	return client
}

// NewData .
func NewData(entClient *ent.Client, logger log.Logger) (*Data, func(), error) {

	logs := log.NewHelper(log.With(logger, "module", "user-service/data"))

	d := &Data{db: entClient}
	return d, func() {
		if err := d.db.Close(); err != nil {
			logs.Error(err)
		}
	}, nil
}
