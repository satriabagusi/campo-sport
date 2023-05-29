package repository

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/midtrans/midtrans-go"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/satriabagusi/campo-sport/internal/entity"
)

type BookingRepository interface {
	GetAllBooking() ([]entity.Booking, error)
	InsertBooking(newBooking *entity.Booking) (*entity.Booking, error)
	GetBookingByOrderNumber(orderNumber string) (*entity.Booking, error)
	UpdateBookingPaymentStatus(orderNumber string) (*entity.Booking, error)
	CancelBooking(orderNumber string) (*entity.Booking, error)
}

type bookingRepository struct {
	db *sql.DB
}

func NewBookingRepository(db *sql.DB) BookingRepository {
	return &bookingRepository{db}
}

func (r *bookingRepository) GetAllBooking() ([]entity.Booking, error) {
	var bookings []entity.Booking
	var voucher_code sql.NullString
	var voucher_discount sql.NullFloat64
	rows, err := r.db.Query(`SELECT b.booking_number, b.total_transaction, 
	u.username, u.phone_number, u.email, 
	c.court_name, c.description, 
	v.voucher_code, v.discount, 
	ts.transaction_status, 
	bd.date_book, bd.start_time, bd.end_time 
	FROM booking_details AS bd 
	INNER JOIN bookings AS b ON bd.booking_id = b.id 
	INNER JOIN users AS u ON b.user_id = u.id 
	INNER JOIN courts AS c ON b.court_id = c.id
	FULL JOIN vouchers AS v ON b.voucher_id = v.id
	INNER JOIN transaction_status AS ts ON b.transaction_status_id = ts.id`)
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var booking entity.Booking
		err := rows.Scan(&booking.BookingNumber, &booking.TotalTransaction, &booking.User.Username, &booking.User.PhoneNumber, &booking.User.Email, &booking.Court.CourtName, &booking.Court.Description, &voucher_code, &voucher_discount, &booking.TransactionStatus.TransactionStatus, &booking.BookingDetail.DateBook, &booking.BookingDetail.StartTime, &booking.BookingDetail.EndTime)
		booking.Voucher.VoucherCode = voucher_code.String
		booking.Voucher.Discount = float32(voucher_discount.Float64)
		if err != nil {
			return nil, err
		}
		bookings = append(bookings, booking)
	}

	return bookings, nil
}

func (r *bookingRepository) InsertBooking(newBooking *entity.Booking) (*entity.Booking, error) {
	//create booking number
	randNumber := rand.Intn(99999)
	year, month, day := time.Now().Date()
	bookingNumber := fmt.Sprintf("%s-%d-%d%d%d", "CMO", randNumber, day, month, year)
	newBooking.CreatedAt = time.Now()
	newBooking.UpdatedAt = time.Now()

	//db query insert booking
	stmt, err := r.db.Prepare(`INSERT INTO bookings (booking_number, user_id, court_id, payment_method_id, voucher_id, total_transaction, transaction_status_id) VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING id`)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	//find court price to calculate total booking
	fCourt, err := r.db.Prepare(`SELECT court_name, court_price FROM courts WHERE id = $1 AND is_available = true`)
	if err != nil {
		log.Println("Error getting court price")
		return nil, err
	}
	defer fCourt.Close()
	fCourtRow := fCourt.QueryRow(newBooking.Court.Id)
	err = fCourtRow.Scan(&newBooking.Court.CourtName, &newBooking.Court.CourtPrice)
	if err != nil {
		return nil, err
	}

	t1 := newBooking.BookingDetail.StartTime
	t2 := newBooking.BookingDetail.EndTime
	totalHour := t2.Sub(t1).Hours()

	newBooking.TotalTransaction = newBooking.Court.CourtPrice * float32(totalHour)

	if newBooking.Voucher.Id != 0 {
		fVoucher, err := r.db.Prepare(`SELECT voucher_code, discount FROM vouchers WHERE id = $1 AND is_available = true`)
		if err != nil {
			log.Println("Error getting court voucher details")
			return nil, err
		}
		defer fVoucher.Close()
		fVoucherRow := fVoucher.QueryRow(newBooking.Voucher.Id)
		err = fVoucherRow.Scan(&newBooking.Voucher.VoucherCode, &newBooking.Voucher.Discount)
		if err != nil {
			return nil, err
		}

		newBooking.TotalTransaction = newBooking.TotalTransaction - newBooking.Voucher.Discount
	}

	if newBooking.PaymentMethod.Id == 2 {
		chargeReq := &coreapi.ChargeReq{
			PaymentType: "bank_transfer",
			TransactionDetails: midtrans.TransactionDetails{
				OrderID:  bookingNumber,
				GrossAmt: int64(newBooking.TotalTransaction),
			},
			BankTransfer: &coreapi.BankTransferDetails{
				Bank: "bca",
			},
		}

		coreApiRes, _ := coreapi.ChargeTransaction(chargeReq)
		newBooking.MidtransResponse = *coreApiRes
	}
	newBooking.TransactionStatus.Id = 1

	err = stmt.QueryRow(bookingNumber, newBooking.User.Id, newBooking.Court.Id, newBooking.PaymentMethod.Id, newBooking.Voucher.Id, newBooking.TotalTransaction, newBooking.TransactionStatus.Id).Scan(&newBooking.Id)

	if err != nil {
		log.Println("Error creating book transaction")
		log.Println(newBooking.Id)
		log.Println(newBooking)
		return nil, err
	}

	rDetailBook, err := r.db.Prepare(`INSERT INTO booking_details (booking_id, date_book, start_time, end_time) VALUES ($1, $2, $3, $4)`)
	if err != nil {

		return nil, err
	}
	defer rDetailBook.Close()
	_, err = rDetailBook.Exec(newBooking.Id, newBooking.BookingDetail.DateBook, newBooking.BookingDetail.StartTime, newBooking.BookingDetail.EndTime)

	if err != nil {
		log.Println("Error creating book details")
		return nil, err
	}

	return newBooking, nil
}

func (r *bookingRepository) GetBookingByOrderNumber(orderNumber string) (*entity.Booking, error) {

	var booking entity.Booking
	var voucher_code sql.NullString
	var voucher_discount sql.NullFloat64

	stmt, err := r.db.Prepare(`SELECT b.booking_number, b.total_transaction, 
	u.username, u.phone_number, u.email, 
	c.court_name, c.description, 
	v.voucher_code, v.discount, 
	ts.transaction_status, 
	bd.date_book, bd.start_time, bd.end_time 
	FROM booking_details AS bd 
	INNER JOIN bookings AS b ON bd.booking_id = b.id 
	INNER JOIN users AS u ON b.user_id = u.id 
	INNER JOIN courts AS c ON b.court_id = c.id
	FULL JOIN vouchers AS v ON b.voucher_id = v.id
	INNER JOIN transaction_status AS ts ON b.transaction_status_id = ts.id
	WHERE b.booking_number = $1`)
	if err != nil {
		return nil, err
	}

	defer stmt.Close()

	row := stmt.QueryRow(orderNumber)
	err = row.Scan(&booking.BookingNumber, &booking.TotalTransaction, &booking.User.Username, &booking.User.PhoneNumber, &booking.User.Email, &booking.Court.CourtName, &booking.Court.Description, &voucher_code, &voucher_discount, &booking.TransactionStatus.TransactionStatus, &booking.BookingDetail.DateBook, &booking.BookingDetail.StartTime, &booking.BookingDetail.EndTime)
	booking.Voucher.VoucherCode = voucher_code.String
	booking.Voucher.Discount = float32(voucher_discount.Float64)
	if err != nil {
		return nil, err
	}

	return &booking, nil
}

func (r *bookingRepository) UpdateBookingPaymentStatus(orderNumber string) (*entity.Booking, error) {
	var booking entity.Booking

	transactionRes, err := coreapi.CheckTransaction(orderNumber)

	if err != nil {
		return nil, err
	}

	booking.MidtransResponse.StatusCode = transactionRes.StatusCode
	booking.MidtransResponse.StatusMessage = transactionRes.StatusMessage
	booking.MidtransResponse.TransactionStatus = transactionRes.TransactionStatus
	booking.MidtransResponse.FraudStatus = transactionRes.FraudStatus
	booking.MidtransResponse.TransactionTime = transactionRes.TransactionTime
	booking.MidtransResponse.Bank = transactionRes.Bank
	booking.MidtransResponse.VaNumbers = transactionRes.VaNumbers

	if transactionRes != nil {
		if transactionRes.TransactionStatus == "pending" {
			fPaymentStatus, err := r.db.Prepare(`SELECT id, transaction_status FROM transaction_status WHERE LOWER(transaction_status) = $1`)
			if err != nil {
				log.Println("SQL Error on get transaction status")
				return nil, err
			}
			defer fPaymentStatus.Close()
			fPaymentStatusRow := fPaymentStatus.QueryRow(transactionRes.TransactionStatus)
			err = fPaymentStatusRow.Scan(&booking.TransactionStatus.Id, &booking.TransactionStatus.TransactionStatus)
			if err != nil {
				log.Println("Error getting transaction status")
				return nil, err
			}

			updateTransaction, err := r.db.Prepare(`UPDATE bookings SET transaction_status_id = $1 WHERE booking_number = $2`)
			if err != nil {

				return nil, err
			}
			defer updateTransaction.Close()

			_, err = updateTransaction.Exec(booking.TransactionStatus.Id, orderNumber)
			if err != nil {
				log.Println("Error update transaction status")
				return nil, err
			}
		} else if transactionRes.TransactionStatus == "settlement" {
			status := "success"
			fPaymentStatus, err := r.db.Prepare(`SELECT id, transaction_status FROM transaction_status WHERE LOWER(transaction_status)=$1 `)
			if err != nil {
				log.Println("SQL Error on get transaction status")
				return nil, err
			}
			defer fPaymentStatus.Close()
			fPaymentStatusRow := fPaymentStatus.QueryRow(status)
			err = fPaymentStatusRow.Scan(&booking.TransactionStatus.Id, &booking.TransactionStatus.TransactionStatus)
			if err != nil {
				log.Println("Error getting transaction status")
				return nil, err
			}

			updateTransaction, err := r.db.Prepare(`UPDATE bookings SET transaction_status_id = $1 WHERE booking_number = $2`)
			if err != nil {
				return nil, err
			}
			defer updateTransaction.Close()

			log.Println(booking.TransactionStatus.Id, booking.TransactionStatus.TransactionStatus)

			_, err = updateTransaction.Exec(booking.TransactionStatus.Id, orderNumber)
			if err != nil {
				log.Println("Error update transaction status")
				return nil, err
			}
		} else if transactionRes.TransactionStatus == "cancel" || transactionRes.TransactionStatus == "expire" {
			fPaymentStatus, err := r.db.Prepare(`SELECT id, transaction_status FROM transaction_status WHERE LOWER(transaction_status) = $1`)
			if err != nil {
				log.Println("SQL Error on get transaction status")
				return nil, err
			}
			defer fPaymentStatus.Close()
			fPaymentStatusRow := fPaymentStatus.QueryRow("failed")
			err = fPaymentStatusRow.Scan(&booking.TransactionStatus.Id, &booking.TransactionStatus.TransactionStatus)
			if err != nil {
				log.Println("Error getting transaction status")
				return nil, err
			}

			updateTransaction, err := r.db.Prepare(`UPDATE bookings SET transaction_status_id = $1 WHERE booking_number = $2`)
			if err != nil {
				return nil, err
			}
			defer updateTransaction.Close()

			_, err = updateTransaction.Exec(booking.TransactionStatus.Id, orderNumber)
			if err != nil {
				log.Println("Error update transaction status")
				return nil, err
			}
		}
	}

	return &booking, nil
}

func (r *bookingRepository) CancelBooking(orderNumber string) (*entity.Booking, error) {

	var booking entity.Booking
	var voucher_code sql.NullString
	var voucher_discount sql.NullFloat64

	stmtTransaction, err := r.db.Prepare(`SELECT b.booking_number, b.total_transaction, 
	u.username, u.phone_number, u.email, 
	c.court_name, c.description, 
	v.voucher_code, v.discount, 
	ts.transaction_status, 
	bd.date_book, bd.start_time, bd.end_time 
	FROM booking_details AS bd 
	INNER JOIN bookings AS b ON bd.booking_id = b.id 
	INNER JOIN users AS u ON b.user_id = u.id 
	INNER JOIN courts AS c ON b.court_id = c.id
	FULL JOIN vouchers AS v ON b.voucher_id = v.id
	INNER JOIN transaction_status AS ts ON b.transaction_status_id = ts.id
	WHERE b.booking_number = $1`)
	if err != nil {
		return nil, err
	}

	defer stmtTransaction.Close()

	row := stmtTransaction.QueryRow(orderNumber)
	err = row.Scan(&booking.BookingNumber, &booking.TotalTransaction, &booking.User.Username, &booking.User.PhoneNumber, &booking.User.Email, &booking.Court.CourtName, &booking.Court.Description, &voucher_code, &voucher_discount, &booking.TransactionStatus.TransactionStatus, &booking.BookingDetail.DateBook, &booking.BookingDetail.StartTime, &booking.BookingDetail.EndTime)
	booking.Voucher.VoucherCode = voucher_code.String
	booking.Voucher.Discount = float32(voucher_discount.Float64)
	if err != nil {
		log.Println("Error while getting detail Booking")
		return nil, err
	}

	if booking.TransactionStatus.TransactionStatus == "Pending" {
		log.Println("When Transaction Status Id is 1 Request Cancel Transaction and Payment Gateway")
		transactionRes, err := coreapi.CancelTransaction(orderNumber)

		if err != nil {
			log.Println("Error while request cancel to payment gateway")
			return nil, err
		}

		booking.MidtransResponse.StatusCode = transactionRes.StatusCode
		booking.MidtransResponse.StatusMessage = transactionRes.StatusMessage
		booking.MidtransResponse.TransactionStatus = transactionRes.TransactionStatus
		booking.MidtransResponse.FraudStatus = transactionRes.FraudStatus
		booking.MidtransResponse.TransactionTime = transactionRes.TransactionTime
		booking.MidtransResponse.Bank = transactionRes.Bank
		booking.MidtransResponse.VaNumbers = transactionRes.VaNumbers

		if transactionRes != nil && transactionRes.TransactionStatus == "cancel" {
			log.Println("When Request Cancel Payment Gateway is approved")

			fPaymentStatus, err := r.db.Prepare(`SELECT id, transaction_status FROM transaction_status WHERE LOWER(transaction_status) = $1`)
			if err != nil {
				log.Println("SQL Error on get transaction status")
				return nil, err
			}
			defer fPaymentStatus.Close()
			fPaymentStatusRow := fPaymentStatus.QueryRow("failed")
			err = fPaymentStatusRow.Scan(&booking.TransactionStatus.Id, &booking.TransactionStatus.TransactionStatus)
			if err != nil {
				log.Println("Error getting transaction status")
				return nil, err
			}

			updateTransaction, err := r.db.Prepare(`UPDATE bookings SET transaction_status_id = $1 WHERE booking_number = $2`)
			if err != nil {
				return nil, err
			}
			defer updateTransaction.Close()

			_, err = updateTransaction.Exec(booking.TransactionStatus.Id, orderNumber)
			if err != nil {
				log.Println("Error update transaction status")
				return nil, err
			}
		}
	} else {
		return nil, errors.New("The Booking status Success/Failed cannot be Cancelled")
	}

	return &booking, nil

}
