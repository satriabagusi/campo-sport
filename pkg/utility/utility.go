package utility

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"golang.org/x/crypto/bcrypt"
)

func GetEnv(key string, v ...any) string {
	//laod .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
	if key != "" {
		return os.Getenv(key)
	}
	return v[0].(string)
}

func Encrypt(str string) string {
	encrptedPassword, err := bcrypt.GenerateFromPassword([]byte(str), bcrypt.DefaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(encrptedPassword)
}
