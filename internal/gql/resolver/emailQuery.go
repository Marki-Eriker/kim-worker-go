package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *queryResolver) Email(ctx context.Context) (*model.EmailQuery, error) {
	return &model.EmailQuery{}, nil
}

// EmailQuery returns runtime.EmailQueryResolver implementation.
func (r *Resolver) EmailQuery() runtime.EmailQueryResolver { return &emailQueryResolver{r} }

type emailQueryResolver struct{ *Resolver }
