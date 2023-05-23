package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type BookingRepository interface {
	GetAllBooking() ([]entity.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) GetAllBooking() ([]entity.Booking, error) {
	var booking []entity.Booking

	return booking, nil
}
