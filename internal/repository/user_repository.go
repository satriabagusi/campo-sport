package repository

import (
	"database/sql"
	"log"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
)

type UserRepository interface {
	UpdateUserStatus(*req.UpdatedStatusUser) (*req.UpdatedStatusUser, error)
	DeleteUser(*entity.User) error
	FindUserById(int) (*res.GetUserByID, error)
	FindUserByEmail(string) (*res.GetUserByUsername, error)
	GetAllUsers() ([]res.GetAllUser, error)
	InsertUser(*req.User) (*res.User, error)
	FindUserByUsername(string) (*res.GetUserByUsername, error)
	FindUserByUsernameLogin(string) (*entity.User, error)
	UpdatePassword(*req.UpdatedPassword) (*req.UpdatedPassword, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) InsertUser(user *req.User) (*res.User, error) {
	stmt, err := r.db.Prepare("INSERT INTO users (username, phone_number, password, email) VALUES ($1, $2, $3, $4) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	err = stmt.QueryRow(user.Username, user.PhoneNumber, user.Password, user.Email).Scan(&user.Id)
	if err != nil {
		return nil, err
	}

	userRes := &res.User{
		Username:    user.Username,
		Email:       user.Email,
		PhoneNumber: user.PhoneNumber,
	}

	return userRes, nil
}

func (r *userRepository) FindUserByUsername(username string) (*res.GetUserByUsername, error) {
	var user res.GetUserByUsername
	stmt, err := r.db.Prepare(`SELECT u.id, u.username, u.phone_number, u.email,  r.role_name, u.is_verified, u.created_at
	FROM users AS u
	JOIN user_roles AS r ON u.role_id = r.id

	WHERE u.username = $1 AND u.is_deleted = false;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	err = row.Scan(&user.Id, &user.Username, &user.PhoneNumber, &user.Email, &user.UserRole, &user.IsVerified, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) FindUserByUsernameLogin(username string) (*entity.User, error) {
	var user entity.User
	stmt, err := r.db.Prepare(`SELECT id, username , password, email, phone_number, created_at, updated_at
	from users
	WHERE username = $1 AND is_deleted = false;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) DeleteUser(user *entity.User) error {
	stmt, err := r.db.Prepare("UPDATE users set is_deleted = true id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(user.Id)
	if err != nil {
		return err
	}

	return nil
}

func (r *userRepository) FindUserById(id int) (*res.GetUserByID, error) {
	var user res.GetUserByID
	stmtm, err := r.db.Prepare(`SELECT u.id, u.username, u.phone_number, u.email, r.role_name, u.is_verified 
	FROM users as u JOIN user_roles as r ON u.role_id = r.id 	
	WHERE u.id = $1 AND u.is_deleted = false`)
	if err != nil {
		return nil, err
	}
	log.Println(id)
	defer stmtm.Close()
	row := stmtm.QueryRow(id)
	err = row.Scan(&user.Id, &user.Username, &user.PhoneNumber, &user.Email, &user.UserRole, &user.IsVerified)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) FindUserByEmail(email string) (*res.GetUserByUsername, error) {
	var user res.GetUserByUsername
	stmtm, err := r.db.Prepare(`SELECT u.id, u.username, u.phone_number, u.email, r.role_name, u.is_verified, u.created_at
	FROM users as u JOIN user_roles as r ON u.role_id = r.id WHERE u.email =$1;`)
	if err != nil {
		return nil, err
	}
	defer stmtm.Close()
	row := stmtm.QueryRow(email)
	err = row.Scan(&user.Id, &user.Username, &user.PhoneNumber, &user.Email, &user.UserRole, &user.IsVerified, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *userRepository) GetAllUsers() ([]res.GetAllUser, error) {
	var users []res.GetAllUser
	rows, err := r.db.Query("SELECT id, username, phone_number, email, is_verified FROM users WHERE is_deleted = false")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user res.GetAllUser
		err := rows.Scan(&user.Id, &user.Username, &user.PhoneNumber, &user.Email, &user.IsVerified)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}

func (r *userRepository) UpdatePassword(updatedUser *req.UpdatedPassword) (*req.UpdatedPassword, error) {
	stmt, err := r.db.Prepare("UPDATE users SET password =$2 WHERE id = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updatedUser.Id, updatedUser.Password)
	if err != nil {
		return nil, err
	}
	return updatedUser, nil
}

func (r *userRepository) UpdateUserStatus(updateStatus *req.UpdatedStatusUser) (*req.UpdatedStatusUser, error) {
	stmt, err := r.db.Prepare("UPDATE users SET role_id =$1 ,is_verified = $2 WHERE id = $3")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(updateStatus.UserRole, updateStatus.IsVerified, updateStatus.Id)
	if err != nil {
		return nil, err
	}
	return updateStatus, nil
}
