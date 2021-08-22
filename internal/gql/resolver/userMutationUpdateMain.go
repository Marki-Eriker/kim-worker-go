package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *userMutationResolver) UpdateMain(ctx context.Context, obj *model.UserMutation, input model.UserUpdateMainInput) (*model.UserUpdateMainOutput, error) {
	u, err := r.app.Services.UserService.UpdateUser(input)
	if err != nil {
		return &model.UserUpdateMainOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserUpdateMainOutput{Ok: true, Record: user.MapOneToGqlModel(u)}, nil
}
