package runtime

import "github.com/pulkitbhardwaj/matrix/internal"

func NewGQLConfig(r Resolver) internal.Config {
	return internal.Config{
		Resolvers:  &r,
		Directives: internal.DirectiveRoot{},
	}
}
