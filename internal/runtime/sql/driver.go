package sql

import (
	"context"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"
)

func NewDriver(lc fx.Lifecycle) (dialect.Driver, error) {
	// Open the database connection.
	drv, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return drv.DB().PingContext(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return drv.Close()
		},
	})

	return drv, nil
}
