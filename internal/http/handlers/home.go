package handlers

import (
	"log/slog"
	"net/http"

	"github.com/axbrunn/gocars/internal/web"
)

type HomeHandler struct {
	logger   *slog.Logger
	renderer *web.Renderer
}

func NewHomeHandler(logger *slog.Logger, renderer *web.Renderer) *HomeHandler {
	return &HomeHandler{
		logger:   logger,
		renderer: renderer,
	}
}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	td := web.TemplateData{Title: "Home"}
	h.renderer.Render(w, http.StatusOK, "home.tmpl", td)
}
