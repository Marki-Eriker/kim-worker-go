package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *paymentMutationResolver) CreateInvoice(ctx context.Context, obj *model.PaymentMutation, input model.PaymentInvoiceCreateInput) (*model.PaymentInvoiceCreateOutput, error) {
	i, err := r.app.Services.PaymentService.CreateInvoice(&input)
	if err != nil {
		return &model.PaymentInvoiceCreateOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.PaymentInvoiceCreateOutput{Ok: true, Record: payment.MapOneInvoiceToGqlModel(i)}, nil
}

func (r *paymentMutationResolver) CreateConfirmation(ctx context.Context, obj *model.PaymentMutation, input model.PaymentConfirmationCreateInput) (*model.PaymentConfirmationCreateOutput, error) {
	c, err := r.app.Services.PaymentService.CreateConfirmation(&input)
	if err != nil {
		return &model.PaymentConfirmationCreateOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.PaymentConfirmationCreateOutput{Ok: true, Record: payment.MapOneConfirmationToGqlModel(c)}, nil
}
