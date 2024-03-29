package req

import "mime/multipart"

type User struct {
	Id          int    `json:"id"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	UserRole    int    `json:"user_role"`
}

type Login struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UpdatedUser struct {
	Id          int    `json:"id"`
	Username    string `json:"username" validate:"required"`
	Password    string `json:"password" validate:"required"`
	Email       string `json:"email" validate:"required"`
	PhoneNumber string `json:"phone_number" validate:"required"`
	UserRole    int    `json:"user_role"`
	IsVerified  bool   `json:"is_verified"`
}

type UpdatedPassword struct {
	Id       int    `json:"id"`
	Password string `json:"password" validate:"required"`
}

type UpdatedStatusUser struct {
	Id         int  `json:"id"`
	UserRole   int  `json:"user_role" validate:"required"`
	IsVerified bool `json:"is_verified" validate:"required"`
}

type UpdateVoucher struct {
	Id          int     `json:"id"`
	VoucherCode string  `json:"voucher_code" validate:"required"`
	IsAvailable bool    `json:"is_available"`
	Discount    float32 `json:"discount" validate:"required,numeric"`
}

type UserDetail struct {
	Id              int     `json:"id"`
	UserId          int     `json:"user_id"`
	Balance         float32 `json:"balance"`
	CredentialProof string  `json:"credential_proof"`
}

type UserCredential struct {
	// Id              int    `json:"id"`
	// UserId          int    `json:"user_id" form:"user_id"`
	Name            string `form:"name"`
	CredentialProof string
	Image           string `form:"image"`
}

type UserProfile struct {
	UserId int                   `form:"user_id" binding:"required"`
	File   *multipart.FileHeader `form:"file" binding:"required"`
}
