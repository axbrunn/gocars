package app

import (
	"log/slog"

	"github.com/axbrunn/gocars/internal/models"
	"github.com/axbrunn/gocars/internal/web"
)

type Application struct {
	Logger    *slog.Logger
	Config    *Config
	Templates web.TemplateCache
	Renderer  *web.Renderer
	Models    models.Models
}
