package runtime

import (
	"go.uber.org/fx"
)

var Module = fx.Module("runtime",
	fx.Provide(
		NewSQLStore,
		// NewGQLConfig,
		NewMuxRouter,
	),
	fx.Invoke(
		NewGQLHandler,
		NewHTTPServer,
	),
)
