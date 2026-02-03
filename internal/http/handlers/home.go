package handlers

import (
	"net/http"

	"github.com/axbrunn/gocars/internal/http/middleware"
	"github.com/axbrunn/gocars/internal/web"
)

type HomeHandler struct {
	renderer *web.Renderer
}

func NewHomeHandler(renderer *web.Renderer) *HomeHandler {
	return &HomeHandler{
		renderer: renderer,
	}
}

func (h *HomeHandler) Index(w http.ResponseWriter, r *http.Request) {
	td := web.NewTemplateData(r)

	theme := "site"
	if tenant, ok := middleware.TenantFromContext(r.Context()); ok {
		td.Title = tenant.Name
		theme = "classic"
	} else {
		td.Title = "GoCars"
	}

	h.renderer.Render(w, r, http.StatusOK, "home.tmpl", theme, td)
}
