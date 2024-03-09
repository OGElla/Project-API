package data

import (
	"database/sql"
	"errors"
)

var (
	ErrRecordNotFound = errors.New("record not found")
)

type Models struct {
	Health HealthModel
}

func NewModels(db *sql.DB) Models{
	return Models{
		Health: HealthModel{DB:db},
	}
}