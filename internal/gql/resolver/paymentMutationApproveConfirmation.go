package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/marki-eriker/kim-worker-go/internal/feature/payment"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *paymentMutationResolver) ApproveConfirmation(ctx context.Context, obj *model.PaymentMutation, input model.PaymentConfirmationApproveInput) (*model.PaymentConfirmationApproveOutput, error) {
	confirm, err := r.app.Services.PaymentService.ApproveConfirmation(&input)
	if err != nil {
		return &model.PaymentConfirmationApproveOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.PaymentConfirmationApproveOutput{Ok: true, Record: payment.MapOneConfirmationToGqlModel(confirm)}, nil
}
