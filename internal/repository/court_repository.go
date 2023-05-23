package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type CourtRepository interface {
	GetAllCourts() ([]entity.Court, error)
}

type courtRepository struct {
	db *sql.DB
}

func NewCourtRepository(db *sql.DB) CourtRepository {
	return &courtRepository{db}
}

func (r *courtRepository) GetAllCourts() ([]entity.Court, error) {
	var user []entity.Court

	return user, nil
}
