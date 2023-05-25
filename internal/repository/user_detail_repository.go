package repository

import (
	"database/sql"
	"fmt"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
)

type UserDetailRepository interface {
	UploadCretential(photo *res.UserProfile) (*res.UserProfile, error)
	CreateUserDetail(*req.UserDetail) error
	GetAllUserDetail() ([]entity.UserDetail, error)
}

type userDetailRepository struct {
	db *sql.DB
}

func NewUserDetailRepository(db *sql.DB) UserDetailRepository {
	return &userDetailRepository{db}
}
func (r *userDetailRepository) CreateUserDetail(userDetail *req.UserDetail) error {
	//Query returning id
	query := `INSERT INTO user_details (user_id, balance)
			VALUES ($1,$2) RETURNING id;`

	err := r.db.QueryRow(query, userDetail.UserId, userDetail.Balance).Scan(&userDetail.Id)

	if err != nil {
		return fmt.Errorf("failed to create user %w", err)
	}
	return nil

}

func (r *userDetailRepository) UploadCretential(userCredential *res.UserProfile) (*res.UserProfile, error) {
	stmt, err := r.db.Prepare("UPDATE user_details SET credential_proof = $1 WHERE user_id = $2")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(userCredential.Url, userCredential.User_id)
	if err != nil {
		return nil, err
	}

	return userCredential, nil
}

func (r *userDetailRepository) GetAllUserDetail() ([]entity.UserDetail, error) {
	var userDetail []entity.UserDetail

	return userDetail, nil
}
