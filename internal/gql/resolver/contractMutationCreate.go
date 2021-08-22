package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *contractMutationResolver) Create(ctx context.Context, obj *model.ContractMutation, input model.ContractCreateInput) (*model.ContractCreateOutput, error) {
	c, err := r.app.Services.ContractService.CreateOrUpdate(&input)
	if err != nil {
		return &model.ContractCreateOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.ContractCreateOutput{Ok: true, Record: contract.MapOneContractToGqlModel(c)}, nil
}
