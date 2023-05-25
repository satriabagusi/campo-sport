package repository

import (
	"database/sql"
	"log"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
)

type CourtRepository interface {
	InsertCourt(*entity.Court) (*res.Court, error)
	UpdateCourt(*entity.Court) (*res.Court, error)
	DeleteCourt(*entity.Court) error
	FindCourtById(int) (*res.Court, error)
	FindCourtByCourt(string) (*res.Court, error)
	GetAllCourts() ([]entity.Court, error)
}

type courtRepository struct {
	db *sql.DB
}

func NewCourtRepository(db *sql.DB) CourtRepository {
	return &courtRepository{db}
}
func (r *courtRepository) InsertCourt(court *entity.Court) (*res.Court, error) {
	stmt, err := r.db.Prepare("INSERT INTO courts (court_name, description, is_available , court_price) VALUES ($1, $2, $3 ,$4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var courtID int
	err = stmt.QueryRow(court.CourtName, court.Description, court.IsAvailable, court.CourtPrice).Scan(&courtID)
	if err != nil {
		return nil, err
	}

	courtRes := &res.Court{
		CourtName:   court.CourtName,
		Description: court.Description,
		CourtPrice:  court.CourtPrice,
		IsAvailable: court.IsAvailable,
	}

	court.Id = courtID

	return courtRes, nil

}

func (r *courtRepository) UpdateCourt(updatedCourt *entity.Court) (*res.Court, error) {
	stmt, err := r.db.Prepare("UPDATE courts SET court_name = $1, description = $2, court_price =$3, is_available = $4 WHERE id = $5")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()
	log.Println("id xourt,", updatedCourt.Id)

	_, err = stmt.Exec(updatedCourt.CourtName, updatedCourt.Description, updatedCourt.CourtPrice, updatedCourt.IsAvailable, updatedCourt.Id)
	if err != nil {
		return nil, err
	}

	courtRes := &res.Court{
		CourtName:   updatedCourt.CourtName,
		Description: updatedCourt.Description,
		CourtPrice:  updatedCourt.CourtPrice,
		IsAvailable: updatedCourt.IsAvailable,
	}
	return courtRes, nil
}

func (r *courtRepository) DeleteCourt(courts *entity.Court) error {
	stmt, err := r.db.Prepare("DELETE FROM courts WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(courts.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *courtRepository) FindCourtById(id int) (*res.Court, error) {
	var court res.Court
	stmtm, err := r.db.Prepare("SELECT court_name, description, court_price, is_available FROM courts WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmtm.Close()
	row := stmtm.QueryRow(id)
	err = row.Scan(&court.CourtName, &court.Description, &court.CourtPrice, &court.IsAvailable)
	if err != nil {
		return nil, err
	}
	return &court, nil
}

func (r *courtRepository) FindCourtByCourt(courtName string) (*res.Court, error) {
	var court res.Court
	stmt, err := r.db.Prepare("SELECT court_name, description, court_price, is_available FROM courts WHERE court_name = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(courtName)
	err = row.Scan(&court.CourtName, &court.Description, &court.CourtPrice, &court.IsAvailable)
	if err != nil {
		return nil, err
	}

	return &court, nil
}

func (r *courtRepository) GetAllCourts() ([]entity.Court, error) {
	var courts []entity.Court

	rows, err := r.db.Query("SELECT id, court_name, description, court_price, is_available FROM courts")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var court entity.Court
		err := rows.Scan(&court.Id, &court.CourtName, &court.Description, &court.CourtPrice, &court.IsAvailable)
		if err != nil {
			return nil, err
		}
		courts = append(courts, court)
	}

	return courts, nil
}
