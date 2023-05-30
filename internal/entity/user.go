/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user.go (c) 2023
Desc: description
Created:  2023-05-22T08:46:11.652Z
Modified: !date!
*/

package entity

import "time"

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	IsVerified  bool      `json:"is_verified"`
	UserRole    int       `json:"user_role"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
