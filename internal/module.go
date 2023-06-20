package internal

import "go.uber.org/fx"

var Module = fx.Module("internal",
	fx.Provide(
		NewGraph,
		NewExecutableSchema,
	),
)
