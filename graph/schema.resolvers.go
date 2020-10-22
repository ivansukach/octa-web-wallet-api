package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/ivansukach/octa-web-wallet-api/graph/generated"
	"github.com/ivansukach/octa-web-wallet-api/graph/model"
	_ "github.com/lib/pq"
)

func (r *queryResolver) Validators(ctx context.Context) ([]*model.Validator, error) {
	return r.validatorRps.Listing(context.Background())
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
