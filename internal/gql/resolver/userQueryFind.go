package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *userQueryResolver) Find(ctx context.Context, obj *model.UserQuery, input model.UserFindInput) (*model.UserFindOutput, error) {
	u, err := r.app.Services.UserService.GetUser(input.UserID)
	if err != nil {
		return &model.UserFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, err
	}

	return &model.UserFindOutput{Ok: true, Record: user.MapOneToGqlModel(&u)}, nil
}
