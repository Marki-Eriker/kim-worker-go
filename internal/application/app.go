package application

import (
	"github.com/marki-eriker/kim-worker-go/internal/database/postgres"
	"go.uber.org/zap"
)

type App struct {
	Logger       *zap.Logger
	Databases    *postgres.Databases
	Repositories *Repositories
	Services     *Services
}

func NewApp(args *Args) (*App, error) {
	logger, err := NewLogger(args.LogLevel)
	if err != nil {
		return nil, err
	}

	databases, err := postgres.NewPostgresDBs(args.PrimaryDatabaseURL, args.LKDatabaseURL, logger)
	if err != nil {
		return nil, err
	}

	repos := NewRepositories(databases)
	services := NewServices(repos)

	app := App{
		Logger:       logger,
		Databases:    databases,
		Repositories: repos,
		Services:     services,
	}

	return &app, nil
}

func (app *App) Close() error {
	if err := app.Databases.Close(); err != nil {
		return err
	}

	return nil
}
