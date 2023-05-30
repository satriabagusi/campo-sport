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

type AdminGetAllUser struct {
	Id              int       `json:"id"`
	Username        string    `json:"username"`
	Email           string    `json:"email"`
	PhoneNumber     string    `json:"phone_number"`
	IsVerified      bool      `json:"is_verified"`
	CreatedAt       time.Time `json:"create_at"`
	CredentialProof any       `json:"credential_proof"`
	UserRole        string    `json:"user_role"`
	IsDeleted       bool      `json:"is_deleted"`
	UpdatedAt       time.Time `json:"updated_at"`
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
	IsVerified bool      `json:"is_verified"`
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

type GetUserProfile struct {
	Id          int        `json:"id"`
	Username    string     `json:"username"`
	Email       string     `json:"email"`
	PhoneNumber string     `json:"phone_number"`
	IsVerified  bool       `json:"is_verified"`
	CreatedAt   time.Time  `json:"create_at"`
	UserRole    int        `json:"role_id"`
	Detail      UserDetail `json:"user_detail"`
	UpdatedAt   time.Time  `json:"updated_at"`
}

type UserDetail struct {
	Balance int    `json:"balance" `
	Url     string `json:"url_credential" `
}

type UserTopUp struct {
	OrderNumber       string    `json:"order_number"`
	Username          string    `json:"username"`
	PaymentMethod     string    `json:"payment_method"`
	Amount            int       `json:"amount"`
	TransactionStatus string    `json:"transaction_status"`
	CreatedAt         time.Time `json:"created_at"`
}

type BookingHistory struct {
	BookingNumber     string  `json:"booking_number"`
	Username          string  `json:"username"`
	CourtName         string  `json:"court_name"`
	TotalTransaction  float64 `json:"total_transaction"`
	VoucherCode       string  `json:"voucher_code"`
	PaymentMethod     string  `json:"payment_method"`
	TransactionStatus string  `json:"transaction_status"`
	CreatedAt         any     `json:"created_at"`
}
