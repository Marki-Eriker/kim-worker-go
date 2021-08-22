package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/feature/file"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
)

func (r *fileMutationResolver) Create(ctx context.Context, obj *model.FileMutation, input model.FileCreateInput) (*model.FileCreateOutput, error) {
	f, err := r.app.Services.FileService.Create(&input)
	if err != nil {
		return &model.FileCreateOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, err
	}

	return &model.FileCreateOutput{Ok: true, Record: file.MapOneToGqlModel(f)}, nil
}
