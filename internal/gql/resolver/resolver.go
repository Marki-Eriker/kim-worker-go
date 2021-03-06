package resolver

import "github.com/marki-eriker/kim-worker-go/internal/application"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	app *application.App
}

func NewResolver(app *application.App) *Resolver {
	return &Resolver{
		app: app,
	}
}
