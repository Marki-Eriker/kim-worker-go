package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
)

func (r *contractQueryResolver) List(ctx context.Context, obj *model.ContractQuery, input model.ContractListInput) (*model.ContactListOutput, error) {
	currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
	if err != nil {
		return &model.ContactListOutput{Ok: false, Error: NewUnauthorizedErrorProblem()}, nil
	}

	contracts, pagination, err := r.app.Services.ContractService.ListContracts(&input, currentUser)
	if err != nil {
		return &model.ContactListOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.ContactListOutput{
		Ok:         true,
		Pagination: pagination,
		Record:     contract.MapManyContractToGqlModel(contracts),
	}, nil
}
