package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/application"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/feature/request"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *contractResolver) ServiceRequest(ctx context.Context, obj *model.Contract) (*model.RequestInfoOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).RequestLoaderByID
	req, err := dataLoader.Load(request.LoaderRequestByIDKey{RequestID: obj.ServiceRequestID})
	if err != nil {
		return &model.RequestInfoOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.RequestInfoOutput{Ok: true, Record: request.MapOneRequestToGqlModel(req)}, nil
}

func (r *contractResolver) Contractor(ctx context.Context, obj *model.Contract) (*model.ContractorGetOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).ContractorLoaderByID
	contractor, err := dataLoader.Load(request.LoaderContractorByIDKey{ContractorID: obj.ContractorID})
	if err != nil {
		return &model.ContractorGetOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.ContractorGetOutput{Ok: true, Record: request.MapOneContractorToGqlModel(contractor)}, nil
}

func (r *contractResolver) PaymentInvoice(ctx context.Context, obj *model.Contract) (*model.PaymentInvoiceFindOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).InvoiceByContractID
	in, err := dataLoader.Load(payment.LoaderInvoiceByContractIDKey{ContractId: obj.ID})
	if err != nil {
		return &model.PaymentInvoiceFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.PaymentInvoiceFindOutput{Ok: true, Record: payment.MapManyInvoiceToGqlModel(in)}, nil
}

// Contract returns runtime.ContractResolver implementation.
func (r *Resolver) Contract() runtime.ContractResolver { return &contractResolver{r} }

type contractResolver struct{ *Resolver }
