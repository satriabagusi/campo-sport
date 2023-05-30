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
	"github.com/midtrans/midtrans-go/iris"
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

type UserWithdraw struct {
	Id                int64                     `json:"id"`
	User              User                      `json:"user"`
	OrderNumber       string                    `json:"order_number"`
	Amount            int                       `json:"amount"`
	PayoutMidtransRes iris.PayoutDetailResponse `json:"withdraw_midtrans_response"`
	BankAccount       string                    `json:"bank_account"`
	BankName          string                    `json:"bank_name"`
	Notes             string                    `json:"notes"`
	CreatedAt         time.Time                 `json:"created_at"`
	UpdatedAt         time.Time                 `json:"updated_at"`
}
