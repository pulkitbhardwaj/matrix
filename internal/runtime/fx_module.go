package runtime

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"go.uber.org/fx"
)

var Module = fx.Module("runtime",
	fx.Provide(
		SQLStore,
		HTTPServer,
		handler.NewDefaultServer,
	),
)
