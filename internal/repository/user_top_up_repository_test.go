/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_top_up_repository_test.go (c) 2023
Desc: description
Created:  2023-05-30T15:31:37.768Z
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

var dummyTopUp = []entity.UserTopUp{
	{
		Id: 1,
		User: entity.User{
			Id:       2,
			Username: "user1",
			Email:    "user1.email@example.com",
		},
		PaymentMethod: entity.PaymentMethod{
			Id: 3,
		},
		OrderNumber: "UTP-182381",
		Amount:      5000000.00,
		MidtransRes: coreapi.ChargeResponse{
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
		TransactionStatus: entity.TransactionStatus{
			Id: 1,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	},
}

type UserTopUpRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *UserTopUpRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}
	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *UserTopUpRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *UserTopUpRepositoryTestSuite) TestUserTopUpRepository_TopUpBalance() {
	var topUp entity.UserTopUp
	topUp.Id = 1

	suite.sqlMock.ExpectPrepare("INSERT INTO user_top_ups").ExpectQuery().WithArgs(dummyTopUp[0].User.Id, dummyTopUp[0].PaymentMethod.Id, dummyTopUp[0].OrderNumber, dummyTopUp[0].Amount, dummyTopUp[0].TransactionStatus.Id, dummyTopUp[0].CreatedAt, dummyTopUp[0].UpdatedAt)

	repo := NewUserTopUpRepository(suite.dbMock)

	u, err := repo.TopUpBalance(&dummyTopUp[0])
	suite.Nil(err)
	suite.Equal(topUp.Id, u.Id)
}

func (suite *UserTopUpRepositoryTestSuite) TestUserTopUpRepository_CheckBalance() {
	panic("implement me")
}

func (suite *UserTopUpRepositoryTestSuite) TestUserTopUpRepository_Withdraw() {
	panic("implement me")
}

func TestUserTopUpRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserTopUpRepositoryTestSuite))
}
