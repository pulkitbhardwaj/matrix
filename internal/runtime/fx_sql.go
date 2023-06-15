package runtime

import (
	"context"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	_ "github.com/mattn/go-sqlite3"
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/matrix/internal"
)

func SQLStore(lc fx.Lifecycle) (*internal.Client, error) {
	// Open the database connection.
	db, err := sql.Open(dialect.SQLite, "file:ent?mode=memory&cache=shared&_fk=1")
	if err != nil {
		return nil, err
	}

	// Create an internal.Client.
	client := internal.NewClient(internal.Driver(entcache.NewDriver(db)))

	lc.Append(fx.Hook{
		// OnStart: func(ctx context.Context) error {
		// 	// Tell the entcache.Driver to skip the caching layer
		// 	// when running the schema migration.
		// 	if client.Schema.Create(entcache.Skip(ctx)); err != nil {
		// 		return err
		// 	}
		// 	return nil
		// },
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
