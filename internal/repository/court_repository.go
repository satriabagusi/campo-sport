/*
Author: Satria Bagus(satria.bagus18@gmail.com)
court_repository.go (c) 2023
Desc: description
Created:  2023-05-23T07:52:18.418Z
Modified: !date!
*/

package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type CourtRepository interface {
	FindCourtById(id int) (*entity.Court, error)
}

type courtRepository struct {
	db *sql.DB
}

func (c *courtRepository) FindCourtById(id int) (*entity.Court, error) {
	var court entity.Court
	stmt, err := c.db.Prepare(`SELECT id, court_name, description, court_price, is_available, created_at, update_at FROM courts WHERE id = $1`)

	if err != nil {
		return nil, err
	}

	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&court.Id, &court.CourtName, &court.Description, &court.CourtPrice, &court.IsAvailable, &court.CreatedAt, &court.UpdatedAt)

	if err != nil {
		return nil, err
	}

	return &court, nil
}

func NewCourtRepository(db *sql.DB) CourtRepository {
	return &courtRepository{db}
}
