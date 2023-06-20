package runtime

import (
	"go.uber.org/fx"

	"github.com/pulkitbhardwaj/matrix/internal/runtime/graphql"
	"github.com/pulkitbhardwaj/matrix/internal/runtime/http"
	"github.com/pulkitbhardwaj/matrix/internal/runtime/sql"
)

var Module = fx.Module("runtime",
	fx.Provide(
		sql.NewDriver,
		graphql.NewConfig,
		http.NewRouter,
	),
	fx.Invoke(
		graphql.Handler,
		http.Server,
	),
)
