package models

import (
	"database/sql"
	"errors"
	"time"

	"github.com/google/uuid"
)

type Tenant struct {
	ID           uuid.UUID
	Slug         string
	Name         string
	TemplateSlug string
	CreatedAt    time.Time
}

type TenantModel struct {
	DB *sql.DB
}

func (m TenantModel) Get(slug string) (*Tenant, error) {
	if len(slug) < 1 {
		return nil, ErrRecordNotFound
	}

	query := `
	SELECT t.id, t.slug, t.name, tpl.slug AS template_slug, t.created_at
	FROM tenants t
	JOIN templates tpl ON t.template_id = tpl.id
	WHERE t.slug = $1
	`

	var tenant Tenant

	err := m.DB.QueryRow(query, slug).Scan(
		&tenant.ID,
		&tenant.Slug,
		&tenant.Name,
		&tenant.TemplateSlug,
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
