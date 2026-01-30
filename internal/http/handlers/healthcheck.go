package handlers

import (
	"fmt"
	"net/http"

	"github.com/axbrunn/gocars/internal/app"
	"github.com/axbrunn/gocars/internal/http/respond"
)

type HealthcheckHandler struct {
	config *app.Config
}

func NewHealthcheckHandler(cfg *app.Config) *HealthcheckHandler {
	return &HealthcheckHandler{
		config: cfg,
	}
}

func (h *HealthcheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprintf(w, "Version: %s\nEnv: %s\n", h.config.Version, h.config.Env)
	if err != nil {
		respond.ServerError(w, r, err)
	}
}
