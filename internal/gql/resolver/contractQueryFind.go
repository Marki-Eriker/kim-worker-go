package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *contractQueryResolver) Find(ctx context.Context, obj *model.ContractQuery, input model.ContractFindInput) (*model.ContractFindOutput, error) {
	c, err := r.app.Services.ContractService.Find(&input)
	if err != nil {
		return &model.ContractFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.ContractFindOutput{Ok: true, Record: contract.MapOneContractToGqlModel(c)}, nil
}
