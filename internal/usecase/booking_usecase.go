package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type BookingUsecase interface {
	GetAllBooking() ([]entity.Booking, error)
}

type bookingUsecase struct {
	bookingRepository repository.BookingRepository
}

func NewBookingUsecase(bookingRepository repository.BookingRepository) BookingUsecase {
	return &bookingUsecase{bookingRepository}
}

func (u *bookingUsecase) GetAllBooking() ([]entity.Booking, error) {
	return u.bookingRepository.GetAllBooking()
}
