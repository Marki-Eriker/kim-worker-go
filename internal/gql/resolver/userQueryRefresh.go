package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *userQueryResolver) Refresh(ctx context.Context, obj *model.UserQuery) (*model.UserRefreshOutput, error) {
	refreshToken, err := middleware.GetRefreshTokenFromCTX(ctx)
	if err != nil {
		return &model.UserRefreshOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	userID, err := r.app.Services.RefreshTokenService.GetTokenUserID(refreshToken)
	if err != nil {
		return &model.UserRefreshOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	user, err := r.app.Services.UserService.GetUser(userID)
	if err != nil {
		return &model.UserRefreshOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	accessToken, err := user.GenToken()
	if err != nil {
		return &model.UserRefreshOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserRefreshOutput{Ok: true, AccessToken: &accessToken}, nil
}
