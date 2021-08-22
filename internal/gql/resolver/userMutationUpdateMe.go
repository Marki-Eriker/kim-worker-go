package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *userMutationResolver) UpdateMe(ctx context.Context, obj *model.UserMutation, input model.UserUpdateMeInput) (*model.UserUpdateMeOutput, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return &model.UserUpdateMeOutput{Ok: false, Error: NewUnauthorizedErrorProblem()}, nil
	}

	if valid, err := input.Validate(); !valid {
		return &model.UserUpdateMeOutput{Ok: false, Error: NewValidationErrorProblem(err)}, nil
	}

	u, err := r.app.Services.UserService.UpdateUserMe(input, currentUser)
	if err != nil {
		return &model.UserUpdateMeOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserUpdateMeOutput{Ok: true, Record: user.MapOneToGqlModel(u)}, nil
}
