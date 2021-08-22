package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *queryResolver) User(ctx context.Context) (*model.UserQuery, error) {
	return &model.UserQuery{ID: "UserQuery"}, nil
}

// UserQuery returns runtime.UserQueryResolver implementation.
func (r *Resolver) UserQuery() runtime.UserQueryResolver { return &userQueryResolver{r} }

type userQueryResolver struct{ *Resolver }
