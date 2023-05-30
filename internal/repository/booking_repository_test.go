/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking_repository_test.go (c) 2023
Desc: description
Created:  2023-05-29T19:48:10.948Z
Modified: !date!
*/

package repository

import (
	"database/sql"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/stretchr/testify/suite"
)

var dummyBooking = []entity.Booking{
	{
		Id:            1,
		BookingNumber: "CMO-0000-2552023",
		PaymentMethod: entity.PaymentMethod{
			Id:            3,
			PaymentMethod: "Online",
		},
		MidtransResponse: coreapi.ChargeResponse{
			StatusCode:        "201",
			StatusMessage:     "Success, Bank Transfer transaction is created",
			TransactionID:     "be03df7d-2f97-4c8c-a53c-8959f1b67295",
			OrderID:           "CMO-0000-2552023",
			GrossAmount:       "450000",
			Currency:          "IDR",
			PaymentType:       "bank_transfer",
			TransactionTime:   time.Now().String(),
			TransactionStatus: "pending",
			VaNumbers: []coreapi.VANumber{
				{
					Bank:     "bca",
					VANumber: "444283823",
				},
			},
			FraudStatus: "accept",
		},
		User: entity.User{
			Id:       2,
			Username: "user1",
			Email:    "user1.email@example.com",
		},
		Court: entity.Court{
			Id:          1,
			CourtName:   "Lapangan Tembak",
			IsAvailable: true,
			Description: "Lapangan khusus latihan menembak Per Jam",
			CourtPrice:  250000,
		},
		TotalTransaction: 450000,
		BookingDetail: entity.BookingDetail{
			Id:        1,
			StartTime: time.Now().Add(time.Hour * 1),
			EndTime:   time.Now().Add(time.Hour * 2),
			BookingId: 1,
			DateBook:  time.Now().Local(),
		},
		Voucher: entity.Voucher{
			Id:          2,
			Discount:    50000,
			IsAvailable: true,
		},
		TransactionStatus: entity.TransactionStatus{
			Id:                1,
			TransactionStatus: "Pending",
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type BookingRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *BookingRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}
	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *BookingRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *BookingRepositoryTestSuite) TestBookingRepository_InsertBooking() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(1324)

	suite.sqlMock.ExpectPrepare("INSERT INTO bookings \\(booking_number, user_id, court_id, payment_method_id, voucher_id, total_transaction, transaction_status_id\\) VALUES \\(\\$1, \\$2, \\$3, \\$4, \\$5, \\$6, \\$7\\) RETURNING id").ExpectQuery().WithArgs(dummyBooking[0].BookingNumber, dummyBooking[0].User.Id, dummyBooking[0].Court.Id, dummyBooking[0].PaymentMethod.Id, dummyBooking[0].Voucher.Id, dummyBooking[0].TotalTransaction, dummyBooking[0].TransactionStatus.Id).WillReturnRows(rows)

	repo := NewBookingRepository(suite.dbMock)

	u, err := repo.InsertBooking(&dummyBooking[0])
	suite.Nil(err)
	suite.Equal(int(1324), u.Id)
}

func TestBookingRepositorySuite(t *testing.T) {
	suite.Run(t, new(BookingRepositoryTestSuite))
}
