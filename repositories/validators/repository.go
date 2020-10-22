package validators

import (
	"context"
	"github.com/ivansukach/octa-web-wallet-api/graph/model"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type validatorRepository struct {
	db *sqlx.DB
}

func New(db *sqlx.DB) Repository {
	return &validatorRepository{db: db}
}

func (r *validatorRepository) Listing(ctx context.Context) ([]*model.Validator, error) {
	rows, err := r.db.Queryx("SELECT * FROM validators")
	if err != nil {
		return nil, err
	}
	validators := make([]*model.Validator, 0)
	for rows.Next() {
		validator := model.Validator{}
		err = rows.StructScan(&validator)
		if err != nil {
			return nil, err
		}
		validators = append(validators, &validator)
	}
	return validators, err
}
