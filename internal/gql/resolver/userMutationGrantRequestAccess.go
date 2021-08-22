package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *userMutationResolver) GrantRequestAccess(ctx context.Context, obj *model.UserMutation, input model.UserGrantRequestAccessInput) (*model.UserGrantRequestAccessOutput, error) {
	u, err := r.app.Services.UserService.UpdateUserServiceTypes(input)
	if err != nil {
		return &model.UserGrantRequestAccessOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserGrantRequestAccessOutput{Ok: true, Record: user.MapOneToGqlModel(u)}, nil
}
