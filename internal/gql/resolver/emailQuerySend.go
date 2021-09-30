package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *emailQueryResolver) Send(ctx context.Context, obj *model.EmailQuery, input model.EmailSendInput) (*model.EmailSendOutput, error) {
	err := r.app.Services.EmailService.SendMail("kyzmin.ig@gmail.com", input.Message)
	if err != nil {
		return &model.EmailSendOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.EmailSendOutput{Ok: true}, nil
}
