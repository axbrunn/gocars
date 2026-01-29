package handlers

import (
	"fmt"
	"log/slog"
	"net/http"

	"github.com/axbrunn/gocars/internal/config"
	"github.com/axbrunn/gocars/internal/http/respond"
)

type HealthcheckHandler struct {
	logger *slog.Logger
	config *config.Config
}

func NewHealthcheckHandler(logger *slog.Logger, cfg *config.Config) *HealthcheckHandler {
	return &HealthcheckHandler{
		config: cfg,
		logger: logger,
	}
}

func (h *HealthcheckHandler) Check(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Server", "Go")
	w.WriteHeader(http.StatusOK)

	_, err := fmt.Fprintf(w, "Version: %s\nEnv: %s\n", h.config.Version, h.config.Env)
	if err != nil {
		respond.ServerError(w, r, h.logger, err)
	}
}
