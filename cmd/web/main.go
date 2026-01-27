package main

import (
	"fmt"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/axbrunn/gocars/internal/config"
	"github.com/axbrunn/gocars/internal/logger"
	"github.com/axbrunn/gocars/internal/router"
)

func main() {
	cfg := config.LoadConfig()
	logger := logger.New()
	slog.SetDefault(logger)

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      router.Routes(cfg),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		ErrorLog:     slog.NewLogLogger(logger.Handler(), slog.LevelError),
	}

	slog.Info("server started", "port", cfg.Port)

	err := srv.ListenAndServe()
	slog.Error(err.Error())
	os.Exit(1)
}
