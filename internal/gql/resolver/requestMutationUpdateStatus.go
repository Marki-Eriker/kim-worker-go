package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *requestMutationResolver) UpdateStatus(ctx context.Context, obj *model.RequestMutation, input model.RequestUpdateStatusInput) (*model.RequestUpdateStatusOutput, error) {
	req, err := r.app.Services.RequestService.UpdateRequestStatus(input)
	if err != nil {
		return &model.RequestUpdateStatusOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.RequestUpdateStatusOutput{Ok: true, Record: request.MapOneRequestToGqlModel(req)}, nil
}
