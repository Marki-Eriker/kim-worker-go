package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *userMutationResolver) Create(ctx context.Context, obj *model.UserMutation, input model.UserCreateInput) (*model.UserCreateOutput, error) {
	if valid, err := input.Validate(); !valid {
		return &model.UserCreateOutput{Ok: false, Error: NewValidationErrorProblem(err)}, nil
	}

	u, err := r.app.Services.UserService.CreateUser(input)
	if err != nil {
		return &model.UserCreateOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserCreateOutput{Ok: true, Record: user.MapOneToGqlModel(u)}, nil
}
