package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"net/http"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *userMutationResolver) Login(ctx context.Context, obj *model.UserMutation, input model.UserLoginInput) (*model.UserLoginOutput, error) {
	w := ctx.Value(middleware.Cookie).(*http.ResponseWriter)
	agent := ctx.Value(middleware.UserAgent).(string)

	accessToken, userID, err := r.app.Services.UserService.Login(input)
	if err != nil {
		return &model.UserLoginOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	err = r.app.Services.RefreshTokenService.CreateRefreshToken(agent, userID, w)
	if err != nil {
		return &model.UserLoginOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.UserLoginOutput{Ok: true, AccessToken: &accessToken}, nil
}
