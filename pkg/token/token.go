package token

import (
	"fmt"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

var (
	mySigningKey     = []byte(utility.GetEnv("SECRET_KEY"))
	expireTimeInt, _ = strconv.Atoi(utility.GetEnv("TOKEN_EXPIRE_TIME_IN_MINUTES"))
)

type MyCustomClaims struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	PhoneNumber string `json:"phone_number"`
	UserRole    int    `json:"user_role" default:"3"`
	IsVerified  bool   `json:"is_verified"`
	jwt.RegisteredClaims
}

func CreateToken(user *entity.User) (string, error) {
	claims := MyCustomClaims{
		user.Id,
		user.Username,
		user.Email,
		user.Password,
		user.PhoneNumber,
		user.UserRole,
		user.IsVerified,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(expireTimeInt) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)

	return ss, err
}

func ValidateToken(tokenString string) (any, error) {
	token, err := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return mySigningKey, nil
	})

	if err != nil {
		return nil, fmt.Errorf("unauthorized")
	}

	claims, ok := token.Claims.(*MyCustomClaims)

	if !ok || !token.Valid {
		return nil, fmt.Errorf("unauthorized")
	}

	return claims, nil
}
