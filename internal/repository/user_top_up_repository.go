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
}

type userTopUpRepository struct {
	db *sql.DB
}

func NewUserTopUpRepository(db *sql.DB) UserTopUpRepository {
	return &userTopUpRepository{db}
}

func (r *userTopUpRepository) TopUpBalance(newTopUp *entity.UserTopUp) (*entity.UserTopUp, error) {
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

	_, err = stmt.Exec(newTopUp.User.Id, newTopUp.PaymentMethod.Id, newTopUp.OrderNumber, newTopUp.Amount, newTopUp.TransactionStatus.Id, newTopUp.CreatedAt, newTopUp.UpdatedAt)

	if err != nil {
		log.Println("Error Creating Top Up Request")
		return nil, err
	}

	return newTopUp, nil
}
