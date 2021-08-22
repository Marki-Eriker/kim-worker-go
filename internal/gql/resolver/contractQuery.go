package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *queryResolver) Contract(ctx context.Context) (*model.ContractQuery, error) {
	return &model.ContractQuery{ID: "ContractQuery"}, nil
}

// ContractQuery returns runtime.ContractQueryResolver implementation.
func (r *Resolver) ContractQuery() runtime.ContractQueryResolver { return &contractQueryResolver{r} }

type contractQueryResolver struct{ *Resolver }
