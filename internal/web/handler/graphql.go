package handler

import (
	"context"
	"errors"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/marki-eriker/kim-worker-go/internal/application"
	"github.com/marki-eriker/kim-worker-go/internal/gql/model"
	"github.com/marki-eriker/kim-worker-go/internal/gql/resolver"
	"github.com/marki-eriker/kim-worker-go/internal/gql/runtime"
	"github.com/marki-eriker/kim-worker-go/internal/web/middleware"
	"github.com/vektah/gqlparser/v2/gqlerror"
	"go.uber.org/zap"
)

func NewGraphQLHandler(app *application.App) *handler.Server {
	h := handler.NewDefaultServer(runtime.NewExecutableSchema(newSchemaConfig(app)))

	h.SetRecoverFunc(func(ctx context.Context, err interface{}) (userMessage error) {
		app.Logger.Error("unhandled error", zap.String("error", fmt.Sprintf("%v", err)))

		return gqlerror.Errorf("internal server error")
	})

	return h
}

func newSchemaConfig(app *application.App) runtime.Config {
	cfg := runtime.Config{
		Resolvers: resolver.NewResolver(app),
	}

	cfg.Directives.HasRole = func(ctx context.Context, obj interface{}, next graphql.Resolver, role model.BaseRole) (interface{}, error) {
		currentUser, err := middleware.GetCurrentUserFromCTX(ctx)
		if err != nil {
			return nil, errors.New("unauthorized")
		}

		if role == model.BaseRoleAny {
			return next(ctx)
		}

		if currentUser.BaseRole != role {
			return nil, errors.New("forbidden")
		}

		return next(ctx)
	}

	return cfg
}
