package req

import "time"

type User struct {
	Id          int       `json:"id"`
	Username    string    `json:"username" validate:"required"`
	Password    string    `json:"password" validate:"required"`
	Email       string    `json:"email" validate:"required"`
	PhoneNumber string    `json:"phone_number" validate:"required"`
	UserRole    int       `json:"user_role"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type UserDetail struct {
	Id              int       `json:"id"`
	UserId          int       `json:"user_id"`
	Balance         float32   `json:"balance"`
	CredentialProof string    `json:"credential_proof"`
	CreatedAt       time.Time `json:"created_at"`
	UpdatedAt       time.Time `json:"updated_at"`
}
