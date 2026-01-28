package main

import (
	"fmt"
	"log/slog"
	"os"
	"time"

	"github.com/axbrunn/gocars/internal/app"
	"github.com/axbrunn/gocars/internal/config"
	"github.com/axbrunn/gocars/internal/logger"
	"github.com/axbrunn/gocars/internal/routes"
	"github.com/axbrunn/gocars/internal/server"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.New()

	app := &app.Application{
		Logger: logger,
		Config: cfg,
	}

	srv := server.NewServer(server.Config{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      routes.SetupRoutes(app),
		Logger:       logger,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	})

	if err := srv.Start(); err != nil {
		slog.Error(err.Error())
		os.Exit(1)
	}
}
