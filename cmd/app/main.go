package main

import (
	"context"
	"fmt"
	"github.com/marki-eriker/kim-worker-go/internal/application"
	"github.com/marki-eriker/kim-worker-go/internal/web/handler"
	"go.uber.org/zap"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	args, err := application.NewArgs()
	if err != nil {
		log.Fatalf("cannot create args: %v", err)
	}

	app, err := application.NewApp(args)
	if err != nil {
		log.Fatalf("cannot create app: %v", err)
	}

	app.Logger.Info("main environment started")

	server := http.Server{
		Addr:    ":" + args.ListenPort,
		Handler: handler.NewAppHandler(app),
	}

	go func() {
		app.Logger.Info(fmt.Sprintf("starting app on port:%v", args.ListenPort))
		if err := server.ListenAndServe(); err != http.ErrServerClosed {
			app.Logger.Fatal("app stopped due error: ", zap.Error(err))
		}

		app.Logger.Info("app stopped gracefully")
	}()

	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	<-interrupt

	app.Logger.Warn("app interruption signal received")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	if err := server.Shutdown(ctx); err != nil {
		app.Logger.Fatal("app shutdown failed", zap.Error(err))
	}

	if err := app.Close(); err != nil {
		app.Logger.Fatal("app environment closing failed", zap.Error(err))
	}
}
