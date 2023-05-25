/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking_usecase.go (c) 2023
Desc: description
Created:  2023-05-23T07:42:05.003Z
Modified: !date!
*/

package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type BookingUsecase interface {
	CreateBooking(newBooking *entity.Booking) (*entity.Booking, error)
}

type bookingUsecase struct {
	bookingRepo repository.BookingRepository
}

func (b *bookingUsecase) CreateBooking(newBooking *entity.Booking) (*entity.Booking, error) {
	return b.bookingRepo.CreateBooking(newBooking)
}
