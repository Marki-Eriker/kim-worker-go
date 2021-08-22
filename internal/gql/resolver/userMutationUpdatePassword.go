package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *userMutationResolver) UpdatePassword(ctx context.Context, obj *model.UserMutation, input model.UserUpdatePasswordInput) (*model.UserUpdatePasswordOutput, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return &model.UserUpdatePasswordOutput{Ok: false, Error: NewUnauthorizedErrorProblem()}, nil
	}

	if currentUser.BaseRole != model.BaseRoleAdmin && currentUser.ID != input.ID {
		return &model.UserUpdatePasswordOutput{Ok: false, Error: NewForbiddenErrorProblem()}, nil
	}

	err = r.app.Services.UserService.UpdatePassword(input)
	if err != nil {
		return &model.UserUpdatePasswordOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserUpdatePasswordOutput{Ok: true}, nil
}
