package middleware

import (
	"context"
	"errors"
	"log/slog"
	"net/http"
	"strings"

	"github.com/axbrunn/gocars/internal/http/respond"
	"github.com/axbrunn/gocars/internal/models"
)

type tenantContexKeyType struct{}

var tenantContextKey = tenantContexKeyType{}

func WithTenant(ctx context.Context, tenant *models.Tenant) context.Context {
	return context.WithValue(ctx, tenantContextKey, tenant)
}

func TenantFromContext(ctx context.Context) (*models.Tenant, bool) {
	tenant, ok := ctx.Value(tenantContextKey).(*models.Tenant)
	return tenant, ok
}

func CheckTenant(m models.Models) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

			host := strings.Split(r.Host, ":")[0]
			parts := strings.Split(host, ".")

			if len(parts) < 2 {
				next.ServeHTTP(w, r)
				return
			}

			slug := parts[0]

			tenant, err := m.Tenants.Get(slug)
			if err != nil {
				switch {
				case errors.Is(err, models.ErrRecordNotFound):
					respond.NotFoundResponse(w, r, err)
				default:
					respond.ServerError(w, r, err)
				}
				return
			}

			slog.Info(
				"tenant resolved",
				"id", tenant.ID,
				"slug", tenant.Slug,
				"name", tenant.Name,
			)

			ctx := WithTenant(r.Context(), tenant)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
