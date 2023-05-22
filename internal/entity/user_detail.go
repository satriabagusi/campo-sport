/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_detail.go (c) 2023
Desc: description
Created:  2023-05-22T08:48:11.527Z
Modified: !date!
*/

package entity

import "time"

type UserDetail struct {
	Id              int       `json:"id"`
	User            User      `json:"user"`
	Balance         float32   `json:"balance"`
	CredentialProof string    `json:"credential_proof"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
