/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking.go (c) 2023
Desc: description
Created:  2023-05-22T08:49:12.322Z
Modified: !date!
*/

package entity

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type Booking struct {
	Id                int                    `json:"id"`
	BookingNumber     string                 `json:"booking_number"`
	User              User                   `json:"user"`
	Court             Court                  `json:"court_detail"`
	PaymentMethod     PaymentMethod          `json:"payment_method"`
	Voucher           Voucher                `json:"voucher_detail"`
	TransactionStatus TransactionStatus      `json:"transaction_status"`
	TotalTransaction  float32                `json:"total_transaction"`
	BookingDetail     BookingDetail          `json:"booking_detail"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
	MidtransResponse  coreapi.ChargeResponse `json:"midtrans_response"`
}
