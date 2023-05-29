package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/stretchr/testify/suite"
)

var dummyCourt = []entity.Court{
	{
		Id:          1,
		CourtName:   "court 1",
		Description: "court 1",
		IsAvailable: true,
		CourtPrice:  200000,
	},
}

type CourtRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *CourtRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}
	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *CourtRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *VoucherRepositoryTestSuite) TestCourtRepository_InsertVoucher() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	//suite.sqlMock.ExpectPrepare("INSERT into courts").ExpectQuery().WithArgs(dummyCourt[0].CourtName, dummyCourt[0].Description, dummyCourt[0].IsAvailable, dummyCourt[0].CourtPrice).WillReturnRows(rows)
	suite.sqlMock.ExpectPrepare("INSERT INTO courts \\(court_name, description, is_available, price\\) VALUES \\(\\$1, \\$2, \\$3, \\$4 \\) RETURNING id").
		ExpectQuery().
		WithArgs(dummyCourt[0].CourtName, dummyCourt[0].Description, dummyCourt[0].IsAvailable, dummyCourt[0].CourtPrice).
		WillReturnRows(rows)

	repo := NewCourtRepository(suite.dbMock)

	u, err := repo.InsertCourt(&dummyCourt[0])
	suite.Nil(err)
	suite.Equal(int(123), u.Id)
}

func (suite *CourtRepositoryTestSuite) TestCourtRepository_FindAllVoucher() {
	rows := sqlmock.NewRows([]string{"id", "court_name", "description", "court_price", "is_available"})
	for _, d := range dummyCourt {
		rows.AddRow(d.Id, d.CourtName, d.Description, d.CourtPrice, d.IsAvailable)
	}
	suite.sqlMock.ExpectQuery("SELECT id, court_name, description, court_price, is_available FROM courts").
		WillReturnRows(rows)

	repo := NewCourtRepository(suite.dbMock)
	users, err := repo.GetAllCourts()
	suite.Nil(err)
	suite.Equal(dummyCourt, users)
}

func TestUserRepositorySuite(t *testing.T) {
	suite.Run(t, new(UserRepositoryTestSuite))
}
