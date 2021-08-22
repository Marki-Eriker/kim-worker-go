package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *contractorResolver) Person(ctx context.Context, obj *model.Contractor) (*model.PersonFindOutput, error) {
	if *obj.PersonID == 0 {
		return &model.PersonFindOutput{Ok: false}, nil
	}

	person, err := r.app.Repositories.RequestRepository.FindPersonByID(ctx, *obj.PersonID)
	if err != nil {
		return &model.PersonFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.PersonFindOutput{Ok: true, Record: request.MapOnePersonToGqlModel(person)}, nil
}

// Contractor returns runtime.ContractorResolver implementation.
func (r *Resolver) Contractor() runtime.ContractorResolver { return &contractorResolver{r} }

type contractorResolver struct{ *Resolver }
