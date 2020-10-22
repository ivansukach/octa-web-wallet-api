package validators

import (
	"context"
	"github.com/ivansukach/octa-web-wallet-api/graph/model"
)

type Repository interface {
	Listing(ctx context.Context) ([]*model.Validator, error)
}
