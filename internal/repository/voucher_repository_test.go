package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/stretchr/testify/suite"
)

var dummyVoucher = []entity.Voucher{
	{
		Id:          1,
		VoucherCode: "voucher 1",
		IsAvailable: true,
		Discount:    10000,
	},
	{
		Id:          2,
		VoucherCode: "voucher 2",
		IsAvailable: true,
		Discount:    20000,
	},
}

type VoucherRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *VoucherRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}
	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *VoucherRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *VoucherRepositoryTestSuite) TestVoucherRepository_InsertVoucher() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	suite.sqlMock.ExpectPrepare("INSERT into vouchers").ExpectQuery().WithArgs(dummyVoucher[0].VoucherCode, dummyVoucher[0].IsAvailable, dummyVoucher[0].Discount).WillReturnRows(rows)

	repo := NewVoucherRepository(suite.dbMock)

	u, err := repo.InsertVoucher(&dummyVoucher[0])
	suite.Nil(err)
	suite.Equal(int(123), u.Id)
}

func TestVoucherRepositorySuite(t *testing.T) {
	suite.Run(t, new(VoucherRepositoryTestSuite))
}
