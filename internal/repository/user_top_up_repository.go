/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_top_up_repository.go (c) 2023
Desc: description
Created:  2023-05-23T11:40:51.701Z
Modified: !date!
*/

package repository

import (
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/satriabagusi/campo-sport/internal/entity"
)

type UserTopUpRepository interface {
	TopUpBalance(newTopUp *entity.UserTopUp) (*entity.UserTopUp, error)
	CheckBalance(orderNumber string) (*entity.UserDetail, error)
	WithdrawBalance(withdrawUser *entity.UserTopUp) (*entity.UserDetail, error)
}

type userTopUpRepository struct {
	db *sql.DB
}

func NewUserTopUpRepository(db *sql.DB) UserTopUpRepository {
	return &userTopUpRepository{db}
}

func (r *userTopUpRepository) TopUpBalance(newTopUp *entity.UserTopUp) (*entity.UserTopUp, error) {

	// BELUM ADA UPDATE SALDO BARU ADA TRANSAKSI KE MIDTRANS
	invoiceNumber := rand.Intn(9999)

	newTopUp.OrderNumber = fmt.Sprintf("%s-%d", "UTP", invoiceNumber)

	stmt, err := r.db.Prepare(`INSERT INTO user_top_ups (user_id, payment_method_id, order_number, amount, transaction_status_id, created_at, updated_at) VALUES ($1, $2, $3, $4, $5, $6, $7)`)

	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	newTopUp.CreatedAt = time.Now()
	newTopUp.UpdatedAt = time.Now()

	findPaymentMethod := r.db.QueryRow(`SELECT payment_method FROM payment_methods WHERE id = $1`, newTopUp.PaymentMethod.Id).Scan(&newTopUp.PaymentMethod.PaymentMethod)

	if findPaymentMethod != nil {
		log.Println("Error find PaymentMethod")
		return nil, err
	}

	findTransactionStatus := r.db.QueryRow(`SELECT transaction_status FROM transaction_status WHERE id = $1`, newTopUp.TransactionStatus.Id).Scan(&newTopUp.TransactionStatus.TransactionStatus)

	if findTransactionStatus != nil {
		log.Println("Error find TransactionStatus")
		log.Println(findTransactionStatus)
		return nil, err
	}

	findUser := r.db.QueryRow(`SELECT username, phone_number, email FROM users WHERE id = $1`, newTopUp.User.Id).Scan(&newTopUp.User.Username, &newTopUp.User.PhoneNumber, &newTopUp.User.Email)

	if findUser != nil {
		log.Println("Error find user or User not Found")
		return nil, err
	}

	if (entity.User{}) != newTopUp.User {
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  newTopUp.OrderNumber,
				GrossAmt: int64(newTopUp.Amount),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: "bca",
			},
		}

		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)

		newTopUp.MidtransRes = *coreApiRes

		_, err := stmt.Exec(newTopUp.User.Id, newTopUp.PaymentMethod.Id, newTopUp.OrderNumber, newTopUp.Amount, newTopUp.TransactionStatus.Id, newTopUp.CreatedAt, newTopUp.UpdatedAt)

		if err != nil {
			log.Println("Error Creating Top Up Request")
			return nil, err
		}
	}

	return newTopUp, nil
}

func (r *userTopUpRepository) CheckBalance(orderNumber string) (*entity.UserDetail, error) {
	var userTopUp entity.UserTopUp
	var userDetail entity.UserDetail

	findUserTopUp, err := r.db.Prepare(`SELECT utu.id, utu.order_number, u.id, u.username, u.phone_number, u.email, utu.amount, pm.payment_method, ts.id, ts.transaction_status, utu.created_at, utu.updated_at 
	FROM user_top_ups AS utu 
	INNER JOIN users AS u ON u.id = utu.user_id 
	INNER JOIN payment_methods AS pm ON pm.id = utu.payment_method_id
	INNER JOIN transaction_status AS ts ON ts.id = utu.transaction_status_id
	WHERE utu.order_number = $1`)

	if err != nil {
		return nil, err
	}
	defer findUserTopUp.Close()
	findUserTopUpRow := findUserTopUp.QueryRow(orderNumber)
	err = findUserTopUpRow.Scan(&userTopUp.Id, &userTopUp.OrderNumber, &userTopUp.User.Id, &userTopUp.User.Username, &userTopUp.User.PhoneNumber, &userTopUp.User.Email, &userTopUp.Amount, &userTopUp.PaymentMethod.PaymentMethod, &userTopUp.TransactionStatus.Id, &userTopUp.TransactionStatus.TransactionStatus, &userTopUp.CreatedAt, &userTopUp.UpdatedAt)
	if err != nil {
		log.Println("Error find User Top Up Detail")
		return nil, err
	}

	transactionRes, erro := coreapi.CheckTransaction(orderNumber)
	if erro != nil {
		log.Println("Error getting transaction detail from midtrans")
		return nil, err
	}

	if transactionRes != nil {
		log.Println("Transaction detail from midtrans returned")
		userTopUp.MidtransRes.StatusCode = transactionRes.StatusCode
		userTopUp.MidtransRes.StatusMessage = transactionRes.StatusMessage
		userTopUp.MidtransRes.TransactionStatus = transactionRes.TransactionStatus
		userTopUp.MidtransRes.FraudStatus = transactionRes.FraudStatus
		userTopUp.MidtransRes.TransactionTime = transactionRes.TransactionTime
		userTopUp.MidtransRes.Bank = transactionRes.Bank
		userTopUp.MidtransRes.VaNumbers = transactionRes.VaNumbers

		if transactionRes.TransactionStatus == "settlement" {
			log.Println("Transaction status settlement")

			getLastBalance := r.db.QueryRow(`SELECT balance FROM user_details WHERE user_id = $1`, userTopUp.User.Id).Scan(&userDetail.Balance)

			log.Println("Last Balance Before Update: ", userDetail.Balance)

			if getLastBalance != nil {
				log.Println("Error find user balance or User Details not Found")
				return nil, err
			}

			log.Println("User Top Up Status :", userTopUp.TransactionStatus.Id)

			var topUpBalance float32
			if userTopUp.TransactionStatus.Id == 1 {
				topUpBalance = userDetail.Balance + float32(userTopUp.Amount)
			} else {
				topUpBalance = userDetail.Balance
			}

			log.Println("Top Up ammount : ", userTopUp.Amount)
			log.Println("Last Balance After Update: ", topUpBalance)
			updateBalance, err := r.db.Prepare(`UPDATE user_details SET balance = $1 WHERE user_id = $2;`)

			if err != nil {
				return nil, err
			}
			defer updateBalance.Close()

			_, err = updateBalance.Exec(topUpBalance, userTopUp.User.Id)

			if err != nil {
				log.Println("Failed to update user balance")
				return nil, err
			}

			status := "success"
			fPaymentStatus, err := r.db.Prepare(`SELECT id, transaction_status FROM transaction_status WHERE LOWER(transaction_status)=$1 `)
			if err != nil {
				log.Println("SQL Error on get transaction status")
				return nil, err
			}
			defer fPaymentStatus.Close()
			fPaymentStatusRow := fPaymentStatus.QueryRow(status)
			err = fPaymentStatusRow.Scan(&userTopUp.TransactionStatus.Id, &userTopUp.TransactionStatus.TransactionStatus)
			if err != nil {
				log.Println("Error getting transaction status")
				return nil, err
			}

			updateTransaction, err := r.db.Prepare(`UPDATE user_top_ups SET transaction_status_id = $1 WHERE order_number = $2`)
			if err != nil {
				return nil, err
			}
			defer updateTransaction.Close()

			log.Println(userTopUp.TransactionStatus.Id, userTopUp.TransactionStatus.TransactionStatus)

			_, err = updateTransaction.Exec(userTopUp.TransactionStatus.Id, orderNumber)
			if err != nil {
				log.Println("Error update transaction status")
				return nil, err
			}

		} else if transactionRes.TransactionStatus == "cancel" || transactionRes.TransactionStatus == "expire" {
			log.Println("Transaction status cancelled or expired")
			fPaymentStatus, err := r.db.Prepare(`SELECT id, transaction_status FROM transaction_status WHERE LOWER(transaction_status) = $1`)
			if err != nil {
				log.Println("SQL Error on get transaction status")
				return nil, err
			}
			defer fPaymentStatus.Close()
			fPaymentStatusRow := fPaymentStatus.QueryRow("failed")
			err = fPaymentStatusRow.Scan(&userTopUp.TransactionStatus.Id, &userTopUp.TransactionStatus.TransactionStatus)
			if err != nil {
				log.Println("Error getting transaction status")
				return nil, err
			}

			updateTransaction, err := r.db.Prepare(`UPDATE user_top_ups SET transaction_status_id = $1 WHERE booking_number = $2`)
			if err != nil {
				return nil, err
			}
			defer updateTransaction.Close()

			_, err = updateTransaction.Exec(userTopUp.TransactionStatus.Id, orderNumber)
			if err != nil {
				log.Println("Error update transaction status")
				return nil, err
			}
		}
	}

	return &userDetail, nil
}

func (r *userTopUpRepository) WithdrawBalance(withdrawUser *entity.UserTopUp) (*entity.UserDetail, error) {
	return nil, nil
}
