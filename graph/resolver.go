package graph

import (
	"github.com/ivansukach/octa-web-wallet-api/repositories/validators"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	validatorRps validators.Repository
}

func NewResolver(validatorRps validators.Repository) *Resolver {
	return &Resolver{validatorRps: validatorRps}
}
