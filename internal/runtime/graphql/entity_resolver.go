package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/pulkitbhardwaj/matrix/internal"
)

func (r *queryResolver) Node(ctx context.Context, id uuid.UUID) (internal.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Nodes(ctx context.Context, ids []uuid.UUID) ([]internal.Noder, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns internal.QueryResolver implementation.
func (r *Resolver) Query() internal.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
