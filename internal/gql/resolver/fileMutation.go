package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *mutationResolver) File(ctx context.Context) (*model.FileMutation, error) {
	return &model.FileMutation{}, nil
}

// FileMutation returns runtime.FileMutationResolver implementation.
func (r *Resolver) FileMutation() runtime.FileMutationResolver { return &fileMutationResolver{r} }

type fileMutationResolver struct{ *Resolver }
