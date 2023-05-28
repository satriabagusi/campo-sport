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
	Username    string    `json:"username"`
	Password    string    `json:"password"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	UserRole    int       `json:"user_role"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
