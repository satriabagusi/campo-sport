/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_detail_repository.go (c) 2023
Desc: description
Created:  2023-05-23T08:19:51.790Z
Modified: !date!
*/
package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type UserDetailRepository interface {
	UpdateBalance(updatedUserDetail *entity.UserDetail) (*entity.UserDetail, error)
}

type userDetailRepository struct {
	db *sql.DB
}

func NewUserDetailRepository(db *sql.DB) UserDetailRepository {
	return &userDetailRepository{db}
}

func (r *userDetailRepository) UpdateBalance(updatedUserDetail *entity.UserDetail) (*entity.UserDetail, error) {



	stmt, err := r.db.Prepare(`UPDATE user_details SET balance = $1 WHERE user_id = $2`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedUserDetail.Balance, updatedUserDetail.User.Id)
	if err != nil {
		return nil, err
	}

	return updatedUserDetail, nil
}
