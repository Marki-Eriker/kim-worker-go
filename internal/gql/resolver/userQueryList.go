package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *userQueryResolver) List(ctx context.Context, obj *model.UserQuery, input model.UserListInput) (*model.UserListOutput, error) {
	users, pagination, err := r.app.Services.UserService.ListUsers(input.Filter)
	if err != nil {
		return &model.UserListOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserListOutput{
		Ok:         true,
		Pagination: pagination,
		Record:     user.MapManyToGqlModels(users),
	}, nil
}
