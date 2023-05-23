package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type UserRepository interface {
	FindUserByUsername(string) (*entity.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
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
