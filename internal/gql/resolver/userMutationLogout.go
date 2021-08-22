package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *userMutationResolver) Logout(ctx context.Context, obj *model.UserMutation) (*model.UserLogoutOutput, error) {
	refreshToken, err := middleware.GetRefreshTokenFromCTX(ctx)
	if err != nil {
		return &model.UserLogoutOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	w, err := middleware.GetResponseWriterFromCTX(ctx)
	if err != nil {
		return &model.UserLogoutOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	err = r.app.Services.RefreshTokenService.Logout(refreshToken, w)
	if err != nil {
		return &model.UserLogoutOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserLogoutOutput{Ok: true}, nil
}
