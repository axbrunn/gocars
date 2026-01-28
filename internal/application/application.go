package application

import (
	"log/slog"

	"github.com/axbrunn/gocars/internal/config"
)

type Application struct {
	Logger *slog.Logger
	Config *config.Config
}
