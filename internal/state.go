package internal

import (
	"context"

	"ariga.io/entcache"
	"entgo.io/ent/dialect"
	"go.uber.org/fx"
)

func NewApplicationState(lc fx.Lifecycle, drv dialect.Driver) *Client {
	// Create a state Client
	state := NewClient(Driver(entcache.NewDriver(drv)))

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			// Tell the entcache.Driver to skip the caching layer
			// when running the schema migration.
			if err := state.Schema.Create(entcache.Skip(ctx)); err != nil {
				return err
			}
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return state.Close()
		},
	})

	return state
}
