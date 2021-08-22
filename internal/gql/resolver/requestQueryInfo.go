package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *requestQueryResolver) Info(ctx context.Context, obj *model.RequestQuery, input model.RequestInfoInput) (*model.RequestInfoOutput, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return &model.RequestInfoOutput{Ok: false, Error: NewUnauthorizedErrorProblem()}, nil
	}

	req, err := r.app.Services.RequestService.GetRequest(input.RequestID, *currentUser)
	if err != nil {
		return &model.RequestInfoOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.RequestInfoOutput{Ok: true, Record: request.MapOneRequestToGqlModel(req)}, nil
}
