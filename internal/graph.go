package internal

import (
	"context"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"go.uber.org/fx"
)

func NewGraph(lc fx.Lifecycle, drv dialect.Driver) *Client {
	client := NewClient(Driver(entcache.NewDriver(drv)))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Tell the entcache.Driver to skip the caching layer
			// when running the schema migration.
			return client.Schema.Create(entcache.Skip(ctx))

		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client
}
