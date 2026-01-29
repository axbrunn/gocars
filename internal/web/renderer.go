package web

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/axbrunn/gocars/internal/http/respond"
)

type Renderer struct {
	templates TemplateCache
	logger    *slog.Logger
}

func NewRenderer(tc TemplateCache, logger *slog.Logger) *Renderer {
	return &Renderer{
		templates: tc,
		logger:    logger,
	}
}

func (r *Renderer) Render(w http.ResponseWriter, status int, page string, td TemplateData) {
	ts, ok := r.templates[page]
	if !ok {
		err := fmt.Errorf("the template %s does not exist", page)
		respond.ServerError(w, nil, r.logger, err)
		return
	}

	w.WriteHeader(status)

	err := ts.ExecuteTemplate(w, "base", td)
	if err != nil {
		respond.ServerError(w, nil, r.logger, err)
	}
}
