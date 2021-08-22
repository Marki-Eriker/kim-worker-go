package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/application"
	"github.com/marki-eriker/kim-worker-go/internal/feature/contract"
	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *requestResolver) ServiceType(ctx context.Context, obj *model.Request) (*model.ServiceTypeGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).ServiceTypeLoaderByID
	types, err := dataLoader.Load(request.LoaderServiceTypesByIDKey{ServiceTypeID: obj.ServiceTypeID})
	if err != nil {
		return &model.ServiceTypeGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.ServiceTypeGetOutput{Ok: true, Record: request.MapOneServiceTypeToGqlModel(types)}, nil
}

func (r *requestResolver) Contractor(ctx context.Context, obj *model.Request) (*model.ContractorGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).ContractorLoaderByID
	contractor, err := dataLoader.Load(request.LoaderContractorByIDKey{ContractorID: obj.ContractorID})
	if err != nil {
		return &model.ContractorGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.ContractorGetOutput{Ok: true, Record: request.MapOneContractorToGqlModel(contractor)}, nil
}

func (r *requestResolver) OrganizationContact(ctx context.Context, obj *model.Request) (*model.OrganizationContactGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).OrganizationContactByID
	contact, err := dataLoader.Load(request.LoaderOrganizationContactByIDKey{OrganizationContactId: *obj.OrganizationContactID})
	if err != nil {
		return &model.OrganizationContactGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.OrganizationContactGetOutput{Ok: true, Record: request.MapOneOrganizationContactToGqlModel(contact)}, nil
}

func (r *requestResolver) BankAccount(ctx context.Context, obj *model.Request) (*model.BankAccountGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).BankAccountLoaderByID
	account, err := dataLoader.Load(request.LoaderBankAccountByIDKey{BankAccountID: *obj.BankAccountID})
	if err != nil {
		return &model.BankAccountGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.BankAccountGetOutput{Ok: true, Record: request.MapOneBankAccountToGqlModel(account)}, nil
}

func (r *requestResolver) Signatory(ctx context.Context, obj *model.Request) (*model.SignatoryGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).SignatoryLoaderByID
	signatory, err := dataLoader.Load(request.LoaderSignatoryByIDKey{SignatoryID: *obj.SignatoryID})
	if err != nil {
		return &model.SignatoryGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.SignatoryGetOutput{Ok: true, Record: request.MapOneSignatoryToGqlModel(signatory)}, nil
}

func (r *requestResolver) Ships(ctx context.Context, obj *model.Request) (*model.ShipGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).ShipLoaderByRequestID
	ships, err := dataLoader.Load(request.LoaderShipByRequestIDKey{RequestID: obj.ID})
	if err != nil {
		return &model.ShipGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.ShipGetOutput{Ok: true, Record: request.MapManyShipToGqlModel(ships)}, nil
}

func (r *requestResolver) Contracts(ctx context.Context, obj *model.Request) (*model.ContactListOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).ContractByRequestID
	contracts, err := dataLoader.Load(contract.LoaderContractByRequestIDKey{RequestID: obj.ID})
	if err != nil {
		return &model.ContactListOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.ContactListOutput{Ok: true, Record: contract.MapManyContractToGqlModel(contracts)}, nil
}

// Request returns runtime.RequestResolver implementation.
func (r *Resolver) Request() runtime.RequestResolver { return &requestResolver{r} }

type requestResolver struct{ *Resolver }
