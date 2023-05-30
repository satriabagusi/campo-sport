/*
Author: Satria Bagus(satria.bagus18@gmail.com)
booking_usecase_test.go (c) 2023
Desc: description
Created:  2023-05-30T16:23:19.620Z
Modified: !date!
*/

package usecase

import (
	"testing"
	"time"

	"github.com/midtrans/midtrans-go/coreapi"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/stretchr/testify/mock"
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

type BookingRepositoryMock struct {
	mock.Mock
}

type BookingUsecaseTestSuite struct {
	suite.Suite
	repoMock *BookingRepositoryMock
}

func (suite *BookingUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(BookingRepositoryMock)
}

func (r *BookingRepositoryMock) GetAllBooking() ([]entity.Booking, error) {
	args := r.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}

	return args.Get(0).([]entity.Booking), nil
}

func(r *BookingRepositoryMock) GetBookingByOrderNumber(string) (*entity.Booking, error) {
	panic("implement me")
}
func(r *BookingRepositoryMock) CreateBooking(newBooking *entity.Booking) (*entity.Booking, error) {
	panic("implement me")
}
func(r *BookingRepositoryMock) UpdateBookingPaymentStatus(string) (*entity.Booking, error) {
	panic("implement me")
}
func(r *BookingRepositoryMock) CancelBooking(string) (*entity.Booking, error) {
	panic("implement me")
}

func TestBookingUsecase(t *testing.T) {
	suite.Run(t, new(BookingUsecaseTestSuite))
}
