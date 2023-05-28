package res

import (
	"time"

	"github.com/google/uuid"
)

type WebResponse struct {
	Code   int         `json:"code"`
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type User struct {
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetAllUser struct {
	Id          int       `json:"id"`
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	IsVerified  bool      `json:"is_verified"`
	CreatedAt   time.Time `json:"create_at"`
	UserRole    int       `json:"role_id"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type GetUserByID struct {
	Id              int    `json:"id"`
	Username        string `json:"username"`
	Email           string `json:"email"`
	PhoneNumber     string `json:"phone_number"`
	CredentialProof any    `json:"credential_proof"`
	IsVerified      string `json:"is_verified"`
	UserRole        string `json:"user_role"`
}
type GetUserByUsername struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	//Password    string    `json:"password"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	PhoneNumber string `json:"phone_number"`
	//CredentialProof string    `json:"credential_proof"`
	IsVerified string    `json:"is_verified"`
	UserRole   string    `json:"user_role"`
	CreatedAt  time.Time `json:"create_at"`
}

type Court struct {
	Id          int     `json:"id"`
	CourtName   string  `json:"court_name"`
	Description string  `json:"description"`
	IsAvailable bool    `json:"is_available"`
	CourtPrice  float32 `json:"courtes_price"`
}

type UserProfile struct {
	User_id int    `json:"user_id" `
	Url     string `json:"url" `
}

type LoginUserResponse struct {
	SessionID             uuid.UUID  `json:"session_id"`
	AccessToken           string     `json:"access_token"`
	AccessTokenExpiresAt  time.Time  `json:"access_token_expires_at"`
	RefreshToken          string     `json:"refresh_token"`
	RefreshTokenExpiresAt time.Time  `json:"refresh_token_expires_at"`
	User                  GetAllUser `json:"user"`
}
