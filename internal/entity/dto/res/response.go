package res

import "time"

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

type Court struct {
	Id          int       `json:"id"`
	CourtName   string    `json:"court_name"`
	Description string    `json:"description"`
	IsAvailable bool      `json:"is_available"`
	CourtPrice  float32   `json:"courtes_price"`
}
