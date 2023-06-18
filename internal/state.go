package internal

import (
	"context"

	"ariga.io/entcache"
	"entgo.io/contrib/entgql"
	"entgo.io/ent/dialect"
	"github.com/google/uuid"
	"go.uber.org/fx"
)

type State interface {
	entgql.TxOpener

	// Noder returns a Node by its id.
	Noder(ctx context.Context, id uuid.UUID, opts ...NodeOption) (Noder, error)

	// Noders returns a Node by its ids.
	Noders(ctx context.Context, ids []uuid.UUID, opts ...NodeOption) ([]Noder, error)
}

func NewApplicationState(lc fx.Lifecycle, drv dialect.Driver) State {
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
