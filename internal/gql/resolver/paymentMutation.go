package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *mutationResolver) Payment(ctx context.Context) (*model.PaymentMutation, error) {
	return &model.PaymentMutation{}, nil
}

// PaymentMutation returns runtime.PaymentMutationResolver implementation.
func (r *Resolver) PaymentMutation() runtime.PaymentMutationResolver {
	return &paymentMutationResolver{r}
}

type paymentMutationResolver struct{ *Resolver }
