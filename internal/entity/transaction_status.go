/*
Author: Satria Bagus(satria.bagus18@gmail.com)
transaction_status.go (c) 2023
Desc: description
Created:  2023-05-22T08:31:37.384Z
Modified: !date!
*/

package entity

import "time"

type TransactionStatus struct {
	Id                int       `json:"id"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
	UpdatedAt         time.Time `json:"updated_at"`
}
