package handlers

import (
	"net/http"

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
	td.Title = "Home"
	h.renderer.Render(w, r, http.StatusOK, "home.tmpl", td)
}
