package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *mutationResolver) Contract(ctx context.Context) (*model.ContractMutation, error) {
	return &model.ContractMutation{}, nil
}

// ContractMutation returns runtime.ContractMutationResolver implementation.
func (r *Resolver) ContractMutation() runtime.ContractMutationResolver {
	return &contractMutationResolver{r}
}

type contractMutationResolver struct{ *Resolver }
