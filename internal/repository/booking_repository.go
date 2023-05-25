/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking_repository.go (c) 2023
Desc: description
Created:  2023-05-23T07:13:24.199Z
Modified: !date!
*/

package repository

import (
	"database/sql"
	"math/rand"
	"strings"
	"time"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type BookingRepository interface {
	CreateBooking(newBooking *entity.Booking) (*entity.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

func (b *bookingRepository) CreateBooking(newBooking *entity.Booking) (*entity.Booking, error) {

	newBooking.BookingNumber = "CMPO" + "000000" + string(rand.Intn(9999)) + strings.Replace(time.DateOnly, "-", "", -1)
	newBooking.CreatedAt = time.Now()
	newBooking.UpdatedAt = time.Now()
	stmt, err := b.db.Prepare(`INSERT INTO bookings (booking_number, user_id, court_id, payment_method_id, voucher_id, total_transaction, transaction_status_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	// totalPrice, err := usecase.NewCourtUsecase()
	// if err != nil {
	// 	return nil, err
	// }

	// newBooking.TotalTransaction = totalPrice

	err = stmt.QueryRow(newBooking.BookingNumber, newBooking.User.Id, newBooking.Court.Id, newBooking.PaymentMethod.Id, newBooking.Voucher.Id, newBooking.TotalTransaction, newBooking.TransactionStatus.Id, newBooking.CreatedAt, newBooking.UpdatedAt).Scan(&newBooking.Id)
	if err != nil {
		return nil, err
	}

	// stmtDetail, err := b.db.Prepare(`INSERT INTO booking_details (booking_id, date_book, start_time, end_time) VALUES ($1, $2, $3, $4)`)
	// if err != nil {
	// 	return nil, err
	// }

	return newBooking, nil
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db}
}
