package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/marki-eriker/kim-worker-go/internal/application"
	"github.com/marki-eriker/kim-worker-go/internal/feature/access"
	"github.com/marki-eriker/kim-worker-go/internal/feature/navigation"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
)

func (r *userResolver) Navigation(ctx context.Context, obj *model.User) (*model.NavigationFindOutput, error) {
	dataLoader := ctx.Value(application.DataLoadersContextKey).(*application.DataLoaders).NavigationLoaderByUserID
	nav, err := dataLoader.Load(navigation.LoaderByUserIDKey{UserID: obj.ID})

	//nav, err := r.app.Repositories.NavigationRepository.GetNavigationForUserID(ctx, obj.ID)
	if err != nil {
		return &model.NavigationFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}
	return &model.NavigationFindOutput{Ok: true, Record: navigation.MapManyToGqlModels(nav)}, nil
}

func (r *userResolver) Access(ctx context.Context, obj *model.User) (*model.AccessFindOutput, error) {
	acc, err := r.app.Repositories.AccessRepository.GetAccessForUser(ctx, obj.ID)
	if err != nil {
		return &model.AccessFindOutput{Ok: false, Error: NewUnknownErrorProblem(err)}, nil
	}

	return &model.AccessFindOutput{Ok: true, Record: access.MapManyToGqlModel(acc)}, nil
}

// User returns runtime.UserResolver implementation.
func (r *Resolver) User() runtime.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
