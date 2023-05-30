/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_top_up.go (c) 2023
Desc: description
Created:  2023-05-22T08:45:36.328Z
Modified: !date!
*/

package entity

import (
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
)

type UserTopUp struct {
	Id                int                    `json:"id"`
	User              User                   `json:"user"`
	PaymentMethod     PaymentMethod          `json:"payment_method"`
	OrderNumber       string                 `json:"order_number"`
	Amount            int                    `json:"amount"`
	MidtransRes       coreapi.ChargeResponse `json:"midtrans_response"`
	TransactionStatus TransactionStatus      `json:"transaction_status"`
	CreatedAt         time.Time              `json:"created_at"`
	UpdatedAt         time.Time              `json:"updated_at"`
}
