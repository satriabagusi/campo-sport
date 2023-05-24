package res

import "time"

type User struct {
	Username    string    `json:"username"`
	Email       string    `json:"email"`
	PhoneNumber string    `json:"phone_number"`
	CreatedAt   time.Time `json:"create_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
