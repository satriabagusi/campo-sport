/*
Author: Satria Bagus(satria.bagus18@gmail.com)
payment_method.go (c) 2023
Desc: description
Created:  2023-05-22T08:30:26.103Z
Modified: !date!
*/

package entity

import "time"

type PaymentMethod struct {
	Id            int       `json:"id"`
	PaymentMethod string    `json:"payment_method"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}
