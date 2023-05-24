package repository

import (
	"database/sql"
	"log"
	"time"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
)

type UserRepository interface {
	InsertUser(*req.User) (*res.User, error)
	FindUserByUsername(string) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user *req.User) (*res.User, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (username, phone_number, password, email, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	user.CreatedAt = time.Now()
	user.UpdatedAt = time.Now()

	err = stmt.QueryRow(user.Username, user.PhoneNumber, user.Password, user.Email, user.CreatedAt, user.UpdatedAt).Scan(&user.Id)
	if err != nil {
		log.Println(user, user.Id)
		log.Println("Failed to insert user:", err)
		return nil, err
	}

	log.Println("User inserted with ID:", user.Id)

	userRes := &res.User{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
		CreatedAt:   user.CreatedAt,
		UpdatedAt:   user.UpdatedAt,
	}

	return userRes, nil
}

func (r *userRepository) FindUserByUsername(username string) (*entity.User, error) {
	var user entity.User
	stmt, err := r.db.Prepare("SELECT id,username,password FROM users WHERE username = $1;")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	err = row.Scan(&user.Id, &user.Username, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
