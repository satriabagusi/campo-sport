/*
Author: Satria Bagus(satria.bagus18@gmail.com)
court.go (c) 2023
Desc: description
Created:  2023-05-22T08:42:43.719Z
Modified: !date!
*/

package entity

import "time"

type Court struct {
	Id          int       `json:"id" `
	CourtName   string    `json:"court_name" validate:"required"`
	Description string    `json:"description" validate:"required"`
	IsAvailable bool      `json:"is_available" `
	CourtPrice  float32   `json:"courtes_price" validate:"required"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
