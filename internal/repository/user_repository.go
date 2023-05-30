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
	AdminGetAllUsers() ([]res.AdminGetAllUser, error)
	InsertUser(*req.User) (*res.User, error)
	FindUserByUsername(string) (*res.GetUserByUsername, error)
	FindUserByUsernameLogin(string) (*entity.User, error)
	FindUserDetailById(int) (res.UserDetail, error)
	UpdatePassword(*req.UpdatedPassword) (*req.UpdatedPassword, error)
	GetAllTopupHistory(id int) ([]res.UserTopUp, error)
	GetAllBookingHistory(id int) ([]res.BookingHistory, error)
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
	stmt, err := r.db.Prepare(`SELECT id, username , password, email, phone_number,role_id, is_verified, created_at, updated_at
	from users
	WHERE username = $1 AND is_deleted = false;`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(username)
	err = row.Scan(&user.Id, &user.Username, &user.Password, &user.Email, &user.PhoneNumber, &user.UserRole, &user.IsVerified, &user.CreatedAt, &user.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (r *userRepository) DeleteUser(user *entity.User) error {
	stmt, err := r.db.Prepare("UPDATE users set is_deleted = true WHERE id = $1")
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

func (r *userRepository) AdminGetAllUsers() ([]res.AdminGetAllUser, error) {
	var users []res.AdminGetAllUser
	rows, err := r.db.Query(`SELECT u.id, u.username, u.phone_number, u.email, r.role_name, u.is_verified , 
	u.is_deleted, ud.credential_proof
	FROM users as u JOIN user_roles as r ON u.role_id = r.id 
	JOIN user_details as ud ON ud.user_id = u.id;`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var user res.AdminGetAllUser
		err := rows.Scan(&user.Id, &user.Username, &user.PhoneNumber, &user.Email, &user.UserRole, &user.IsVerified, &user.IsDeleted, &user.CredentialProof)
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

func (r *userRepository) FindUserDetailById(id int) (res.UserDetail, error) {
	var user res.UserDetail
	stmtm, err := r.db.Prepare(`SELECT credential_proof, balance 
	FROM user_details JOIN users ON user_details.user_id = users.id 	
	WHERE users.id = $1;`)
	if err != nil {
		return user, err
	}
	log.Println(id)
	defer stmtm.Close()
	row := stmtm.QueryRow(id)
	err = row.Scan(&user.Url, &user.Balance)
	if err != nil {
		return user, err
	}
	return user, nil
}

func (r *userRepository) GetAllTopupHistory(id int) ([]res.UserTopUp, error) {
	var users []res.UserTopUp
	rows, err := r.db.Query(`Select t.order_number, u.username, t.amount,  ts.transaction_status, t.created_at, m.payment_method 
	From user_top_ups as t JOIN users as u on u.id = t.user_id
	Join transaction_status as ts on ts.id = t.transaction_status_id 
	join payment_methods as m on m.id = t.payment_method_id
	where t.user_id = $1;`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// err = row.Scan(&user.OrderNumber, &user.Username, &user.Amount, &user.TransactionStatus, &user.CreatedAt, &user.PaymentMethod)
	for rows.Next() {
		var user res.UserTopUp
		err := rows.Scan(&user.OrderNumber, &user.Username, &user.Amount, &user.TransactionStatus, &user.CreatedAt, &user.PaymentMethod)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}

func (r *userRepository) GetAllBookingHistory(id int) ([]res.BookingHistory, error) {
	var users []res.BookingHistory
	rows, err := r.db.Query(`SELECT b.booking_number, u.username, b.total_transaction, v.voucher_code ,  ts.transaction_status, 
	c.court_name , m.payment_method, b.created_at
	From bookings as b JOIN users as u on u.id = b.user_id
	Join transaction_status as ts on ts.id = b.transaction_status_id 
	join payment_methods as m on m.id = b.payment_method_id
	Join courts as c on c.id = b.court_id
	Join vouchers as v on v.id = b.voucher_id
	where b.user_id = $1;`, id)
	if err != nil {
		return nil, err
	}

	defer rows.Close()
	// err = row.Scan(&user.OrderNumber, &user.Username, &user.Amount, &user.TransactionStatus, &user.CreatedAt, &user.PaymentMethod)
	for rows.Next() {
		var user res.BookingHistory
		err := rows.Scan(&user.BookingNumber, &user.Username, &user.TotalTransaction, &user.VoucherCode, &user.TransactionStatus, &user.CourtName, &user.PaymentMethod, &user.CreatedAt)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	if err != nil {
		return nil, err
	}
	return users, nil
}
