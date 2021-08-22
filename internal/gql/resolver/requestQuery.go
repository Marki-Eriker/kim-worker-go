package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *queryResolver) Request(ctx context.Context) (*model.RequestQuery, error) {
	return &model.RequestQuery{ID: "RequestQuery"}, nil
}

// RequestQuery returns runtime.RequestQueryResolver implementation.
func (r *Resolver) RequestQuery() runtime.RequestQueryResolver { return &requestQueryResolver{r} }

type requestQueryResolver struct{ *Resolver }
