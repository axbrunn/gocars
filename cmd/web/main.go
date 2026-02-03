package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/axbrunn/gocars/internal/app"
	"github.com/axbrunn/gocars/internal/http/routes"
	"github.com/axbrunn/gocars/internal/logger"
	"github.com/axbrunn/gocars/internal/models"
	"github.com/axbrunn/gocars/internal/server"
	"github.com/axbrunn/gocars/internal/web"
)

func main() {
	cfg := app.LoadConfig()
	logger := logger.New()
	slog.SetDefault(logger)

	logger.Info("connecting to database")

	db, err := app.OpenDB(*cfg)
	if err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}

	defer func() {
		logger.Info("closing database connection")
		db.Close()
	}()

	logger.Info("database connection pool established")

	templateCache, err := web.NewTemplateCache()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}

	renderer := web.NewRenderer(templateCache)

	app := &app.Application{
		Logger:    logger,
		Config:    cfg,
		Templates: templateCache,
		Renderer:  renderer,
		Models:    models.NewModels(db),
	}

	srv := server.NewServer(server.Config{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      routes.SetupRoutes(app),
		Logger:       logger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	err = srv.Start()
	if err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
