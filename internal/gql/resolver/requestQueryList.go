package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *requestQueryResolver) List(ctx context.Context, obj *model.RequestQuery, input model.RequestListInput) (*model.RequestListOutput, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return &model.RequestListOutput{Ok: false, Error: NewUnauthorizedErrorProblem()}, nil
	}

	requests, pagination, err := r.app.Services.RequestService.ListRequests(&input, currentUser)
	if err != nil {
		return &model.RequestListOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.RequestListOutput{
		Ok:         true,
		Pagination: pagination,
		Record:     request.MapManyRequestToGqlModels(requests),
	}, nil
}
