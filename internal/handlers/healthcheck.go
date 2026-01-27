package handlers

import (
	"fmt"
	"net/http"

	"github.com/axbrunn/gocars/internal/config"
)

func HealthcheckHandler(cfg config.Config) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Server", "Go")
		w.WriteHeader(http.StatusOK)

		fmt.Fprintf(w, "Version: %s\nEnv: %s\n", cfg.Version, cfg.Env)
	}
}
