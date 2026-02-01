package middleware

import (
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/axbrunn/gocars/internal/http/respond"
	"github.com/axbrunn/gocars/internal/models"
)

func CheckTenant(m models.Models) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			host := strings.Split(r.Host, ":")[0]
			parts := strings.Split(host, ".")

			// TODO: Should redirect to localhost without tenant
			if len(parts) < 2 {
				http.Error(w, "tenant is missing", http.StatusBadRequest)
				return
			}

			slug := parts[0]

			tenant, err := m.Tenants.Get(slug)
			if err != nil {
				switch {
				// TODO: Should redirect to localhost without tenant
				case errors.Is(err, models.ErrRecordNotFound):
					respond.NotFoundResponse(w, r, err)
				default:
					respond.ServerError(w, r, err)
				}
				return
			}

			slog.Info("received request from", "ID", tenant.ID, "name", tenant.Name)

			next.ServeHTTP(w, r)
		})
	}
}
