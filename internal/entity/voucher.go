/*
Author: Satria Bagus(satria.bagus18@gmail.com)
voucher.go (c) 2023
Desc: description
Created:  2023-05-22T08:38:05.593Z
Modified: !date!
*/

package entity

import "time"

type Voucher struct {
	Id          int       `json:"id"`
	VoucherCode string    `json:"voucher_code"`
	IsAvailable bool      `json:"is_available"`
	Discount    float32   `json:"discount"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
