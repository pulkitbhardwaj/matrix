package runtime

import (
	"github.com/pulkitbhardwaj/matrix/internal"
	"go.uber.org/fx"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	fx.In

	State internal.State
}
