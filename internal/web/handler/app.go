package handler

import (
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/marki-eriker/kim-worker-go/internal/application"
	customMiddleware "github.com/marki-eriker/kim-worker-go/internal/web/middleware"
	"github.com/rs/cors"
)

func NewAppHandler(app *application.App) *chi.Mux {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Use(customMiddleware.AuthMiddleware(app.Repositories.UserRepository))
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3010", "http://localhost:5000", "http://10.1.100.120:3010"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
		Debug:            false,
	}).Handler)

	dataLoaderInjector := customMiddleware.NewDataLoadersMiddleware(app)

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", dataLoaderInjector(NewGraphQLHandler(app)))
	router.Post("/upload", uploadHandler())
	//router.Handle("/query", NewGraphQLHandler(app))

	return router
}
