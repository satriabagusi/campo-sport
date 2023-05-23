package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type UserDetailRepository interface {
	GetAllUserDetail() ([]entity.UserDetail, error)
}

type userDetailRepository struct {
	db *sql.DB
}

func NewUserDetailRepository(db *sql.DB) UserDetailRepository {
	return &userDetailRepository{db}
}

func (r *userDetailRepository) GetAllUserDetail() ([]entity.UserDetail, error) {
	var userDetail []entity.UserDetail

	return userDetail, nil
}
