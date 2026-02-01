package models

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Tenants TenantModel
}

func NewModels(db *sql.DB) Models {
	return Models{
		Tenants: TenantModel{DB: db},
	}
}
