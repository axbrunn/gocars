package handlers

import (
	"net/http"

	"github.com/axbrunn/gocars/internal/web"
)

type VehicleHandler struct {
	renderer *web.Renderer
}

func NewVehicleHandler(renderer *web.Renderer) *VehicleHandler {
	return &VehicleHandler{
		renderer: renderer,
	}
}

func (h *VehicleHandler) View(w http.ResponseWriter, r *http.Request) {
	td := web.NewTemplateData(r)

	h.renderer.Render(w, r, http.StatusOK, "vehicle.tmpl", td)
}

func (h *VehicleHandler) List(w http.ResponseWriter, r *http.Request) {
	td := web.NewTemplateData(r)

	h.renderer.Render(w, r, http.StatusOK, "vehicles.tmpl", td)
}
