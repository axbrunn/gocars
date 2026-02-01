package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID        uuid.UUID
	Slug      string
	Name      string
	CreatedAt time.Time
}

type TenantModel struct {
	DB *sql.DB
}

func (m TenantModel) Get(slug string) (*Tenant, error) {
	if len(slug) < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
        SELECT id, slug, name, created_at
        FROM tenants
        WHERE slug = $1`

	var tenant Tenant

	err := m.DB.QueryRow(query, slug).Scan(
		&tenant.ID,
		&tenant.Slug,
		&tenant.Name,
		&tenant.CreatedAt,
	)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &tenant, nil
}
