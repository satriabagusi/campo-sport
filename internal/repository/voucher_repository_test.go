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
	suite.sqlMock.ExpectPrepare("INSERT INTO vouchers \\(voucher_code, is_available, discount\\) VALUES \\(\\$1, \\$2, \\$3\\) RETURNING id").
		ExpectQuery().
		WithArgs(dummyVoucher[0].VoucherCode, dummyVoucher[0].IsAvailable, dummyVoucher[0].Discount).
		WillReturnRows(rows)
		
	repo := NewVoucherRepository(suite.dbMock)

	u, err := repo.InsertVoucher(&dummyVoucher[0])
	suite.Nil(err)
	suite.Equal(int(123), u.Id)
}

// func (suite *VoucherRepositoryTestSuite) TestVoucherRepository_UpdateVoucher() {
// 	// Persiapan input dan output yang diharapkan
// 	voucher := &req.UpdateVoucher{
// 		Id:          1,
// 		VoucherCode: "UPDATED",
// 		IsAvailable: false,
// 		Discount:    5000,
// 	}

// 	// Mengatur ekspektasi query UPDATE
// 	suite.sqlMock.ExpectPrepareRegex("UPDATE vouchers SET voucher_code = (.+), is_available = (.+), discount = (.+) WHERE id = (.+)").
// 		ExpectQuery().
// 		WithArgs(voucher.VoucherCode, voucher.IsAvailable, voucher.Discount, voucher.Id).
// 		WillReturnResult(sqlmock.NewResult(0, 1))

// 	// Membuat objek voucherRepository dengan database mock
// 	repo := voucherRepository{db: suite.dbMock}

// 	// Memanggil fungsi yang akan diuji
// 	result, err := repo.UpdateVoucher(voucher)

// 	// Memastikan tidak ada error saat menjalankan query
// 	suite.NoError(err)

// 	// Memastikan hasil sesuai dengan yang diharapkan
// 	suite.Equal(voucher, result)

// 	// Memastikan semua ekspektasi query terpenuhi
// 	err = suite.sqlMock.ExpectationsWereMet()
// 	suite.NoError(err)
// }

func (suite *VoucherRepositoryTestSuite) TestVoucherRepository_FindAllVoucher() {
	rows := sqlmock.NewRows([]string{"id", "voucher_code", "is_available", "discount"})
	for _, d := range dummyVoucher {
		rows.AddRow(d.Id, d.VoucherCode, d.IsAvailable, d.Discount)
	}
	suite.sqlMock.ExpectQuery("SELECT id, voucher_code, is_available, discount FROM vouchers").
		WillReturnRows(rows)

	repo := NewVoucherRepository(suite.dbMock)
	users, err := repo.GetAllVoucher()
	suite.Nil(err)
	suite.Equal(dummyVoucher, users)
}

// func (suite *VoucherRepositoryTestSuite) TestUpdateVoucher_Success() {
// 	// Membuat objek voucherRepository dengan database mock
// 	repo := voucherRepository{db: suite.dbMock}

// 	// Persiapan input dan output yang diharapkan
// 	voucher := &req.UpdateVoucher{
// 		Id:          1,
// 		VoucherCode: "UPDATED",
// 		IsAvailable: false,
// 		Discount:    5000,
// 	}
// 	expectedVoucher := &req.UpdateVoucher{
// 		Id:          1,
// 		VoucherCode: "UPDATED",
// 		IsAvailable: false,
// 		Discount:    5000,
// 	}

// 	// Mengatur ekspektasi query UPDATE
// 	suite.sqlMock.ExpectPrepare("UPDATE vouchers SET voucher_code = (.+), is_available = (.+), discount = (.+) WHERE id = (.+)")
// 	suite.sqlMock.ExpectExec().WithArgs(voucher.VoucherCode, voucher.IsAvailable, voucher.Discount, voucher.Id).WillReturnResult(sqlmock.NewResult(0, 1))

// 	// Memanggil fungsi yang akan diuji
// 	result, err := repo.UpdateVoucher(voucher)

// 	// Memastikan tidak ada error saat menjalankan query
// 	suite.NoError(err)

// 	// Memastikan hasil sesuai dengan yang diharapkan
// 	suite.Equal(expectedVoucher, result)

// 	// Memastikan semua ekspektasi query terpenuhi
// 	err = suite.sqlMock.ExpectationsWereMet()
// 	suite.NoError(err)
// }

func TestVoucherRepositorySuite(t *testing.T) {
	suite.Run(t, new(VoucherRepositoryTestSuite))
}
