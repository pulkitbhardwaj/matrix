package flow

import (
	"context"

	"github.com/onflow/flow-go-sdk/access"
	"github.com/onflow/flow-go-sdk/access/http"
	"go.uber.org/fx"
)

func NewClient(lc fx.Lifecycle) (access.Client, error) {
	client, err := http.NewClient(http.EmulatorHost)
	if err != nil {
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			return client.Ping(ctx)
		},
		OnStop: func(ctx context.Context) error {
			return client.Close()
		},
	})

	return client, nil
}
