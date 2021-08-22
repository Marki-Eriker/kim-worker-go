package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *mutationResolver) Request(ctx context.Context) (*model.RequestMutation, error) {
	return &model.RequestMutation{}, nil
}

// RequestMutation returns runtime.RequestMutationResolver implementation.
func (r *Resolver) RequestMutation() runtime.RequestMutationResolver {
	return &requestMutationResolver{r}
}

type requestMutationResolver struct{ *Resolver }
