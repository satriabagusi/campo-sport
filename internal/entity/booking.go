/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking.go (c) 2023
Desc: description
Created:  2023-05-22T08:49:12.322Z
Modified: !date!
*/

package entity

import "time"

type Booking struct {
	Id                int               `json:"id"`
	BookingNumber     string            `json:"booking_number"`
	Court             Court             `json:"court_detail"`
	PaymentMethod     PaymentMethod     `json:"payment_method"`
	Voucher           Voucher           `json:"voucher_detail"`
	TransactionStatus TransactionStatus `json:"transaction_status"`
	CreatedAt         time.Time         `json:"created_at"`
	UpdatedAt         time.Time         `json:"updated_at"`
}
