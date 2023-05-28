package token

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

// var (
// 	secretKey         = utility.GetEnv("SECRET_KEY")
// 	expireTmeInt, err = strconv.Atoi(utility.GetEnv("TOKEN_EXPIRE_TIME_IN_MINUTES"))
// )

type Token struct {
	UserID    int
	ExpiresAt int64
}

func GenerateToken(userID int) (string, error) {
	//expireTimeInt, _ := strconv.Atoi(utility.GetEnv("TOKEN_EXPIRE_TIME_IN_MINUTES"))
	expiresAt := time.Now().Add(time.Duration(1) * time.Minute)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userID":    userID,
		"expiresAt": expiresAt,
	})

	secretKey := []byte(utility.GetEnv("SECRET_KEY"))

	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ParseToken(tokenString string, secretKey []byte) (*Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Make sure the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	userID := int(claims["userID"].(float64))
	expiresAt := int64(claims["expiresAt"].(float64))

	return &Token{
		UserID:    userID,
		ExpiresAt: expiresAt,
	}, nil
}

// type Token struct {
// 	UserID    int64
// 	ExpiresAt int64
// }

// func GenerateToken(userID int64, expiresAt time.Time, secretKey []byte) (string, error) {
// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
// 		"userID":    userID,
// 		"expiresAt": expiresAt.Unix(),
// 	})

// 	tokenString, err := token.SignedString(secretKey)
// 	if err != nil {
// 		return "", err
// 	}

// 	return tokenString, nil
// }

// func ParseToken(tokenString string, secretKey []byte) (*Token, error) {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		// Make sure the signing method is correct
// 		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// 			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// 		}

// 		return secretKey, nil
// 	})

// 	if err != nil {
// 		return nil, err
// 	}

// 	claims, ok := token.Claims.(jwt.MapClaims)
// 	if !ok || !token.Valid {
// 		return nil, fmt.Errorf("invalid token")
// 	}

// 	userID := int64(claims["userID"].(float64))
// 	expiresAt := int64(claims["expiresAt"].(float64))

// 	return &Token{
// 		UserID:    userID,
// 		ExpiresAt: expiresAt,
// 	}, nil
// }
