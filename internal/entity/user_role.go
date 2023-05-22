/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_role.go (c) 2023
Desc: description
Created:  2023-05-22T08:29:19.244Z
Modified: !date!
*/

package entity

import "time"

type UserRole struct {
	Id        int       `json:"id"`
	RoleName  string    `json:"role_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
