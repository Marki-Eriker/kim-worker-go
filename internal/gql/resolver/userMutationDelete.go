package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *userMutationResolver) Delete(ctx context.Context, obj *model.UserMutation, id uint) (*model.UserDeleteOutput, error) {
	err := r.app.Services.UserService.DeleteUser(id)
	if err != nil {
		return &model.UserDeleteOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserDeleteOutput{Ok: true}, nil
}
