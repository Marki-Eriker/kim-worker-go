package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/user"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *userQueryResolver) Me(ctx context.Context, obj *model.UserQuery) (*model.UserMeOutput, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return &model.UserMeOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserMeOutput{Ok: true, Record: user.MapOneToGqlModel(currentUser)}, nil
}
