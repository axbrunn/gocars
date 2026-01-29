package handlers

import (
	"html/template"
	"log/slog"
	"net/http"

	"github.com/axbrunn/gocars/internal/http/respond"
)

type HomeHandler struct {
	logger *slog.Logger
}

func NewHomeHandler(logger *slog.Logger) *HomeHandler {
	return &HomeHandler{logger: logger}
}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")

	files := []string{
		"internal/templates/layout/base.tmpl",
		"internal/templates/layout/nav.tmpl",
		"internal/templates/public/home.tmpl",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		respond.ServerError(w, r, h.logger, err)
		return
	}

	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		respond.ServerError(w, r, h.logger, err)
	}
}
