package middleware

import (
	"context"
	"net/http"

	"github.com/marki-eriker/kim-worker-go/internal/application"
)

func NewDataLoadersMiddleware(app *application.App) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			ctx := context.WithValue(r.Context(), application.DataLoadersContextKey, application.NewDataLoaders(app.Repositories))
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
