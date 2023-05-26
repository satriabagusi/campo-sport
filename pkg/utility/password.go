package utility

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

// HashPassword returns the bycrypt hash of the password
func HashPassword(password string) (string, error) {
	hashsedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", fmt.Errorf("failed to hash password: %w", err)

	}
	return string(hashsedPassword), nil
}

// CheckPassowrd checks id the provided password is correct or not
func CheckPassword(password string, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
