package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/application"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *paymentInvoiceResolver) Confirmation(ctx context.Context, obj *model.PaymentInvoice) (*model.PaymentConfirmationFindOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).ConfirmationByInvoiceID
	confirm, err := dataLoader.Load(payment.LoaderConfirmationByInvoiceIDKey{InvoiceID: obj.ID})
	if err != nil {
		return &model.PaymentConfirmationFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.PaymentConfirmationFindOutput{Ok: true, Record: payment.MapOneConfirmationToGqlModel(confirm)}, nil
}

// PaymentInvoice returns runtime.PaymentInvoiceResolver implementation.
func (r *Resolver) PaymentInvoice() runtime.PaymentInvoiceResolver { return &paymentInvoiceResolver{r} }

type paymentInvoiceResolver struct{ *Resolver }
