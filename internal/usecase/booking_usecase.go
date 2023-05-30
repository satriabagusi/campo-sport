package usecase

import (
	"errors"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type BookingUsecase interface {
	GetAllBooking() ([]entity.Booking, error)
	GetBookingByOrderNumber(string) (*entity.Booking, error)
	CreateBooking(newBooking *entity.Booking) (*entity.Booking, error)
	UpdateBookingPaymentStatus(string) (*entity.Booking, error)
	CancelBooking(string) (*entity.Booking, error)
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

func (u *bookingUsecase) CreateBooking(newBooking *entity.Booking) (*entity.Booking, error) {
	if newBooking.Court.Id == 0 {
		return nil, errors.New("Please choose court to book")
	}

	return u.bookingRepository.InsertBooking(newBooking)
}

func (u *bookingUsecase) GetBookingByOrderNumber(bookingNumber string) (*entity.Booking, error) {
	return u.bookingRepository.GetBookingByOrderNumber(bookingNumber)
}

func (u *bookingUsecase) UpdateBookingPaymentStatus(orderNumber string) (*entity.Booking, error) {
	return u.bookingRepository.UpdateBookingPaymentStatus(orderNumber)
}

func (u *bookingUsecase) CancelBooking(orderNumber string) (*entity.Booking, error) {
	return u.bookingRepository.CancelBooking(orderNumber)
}
