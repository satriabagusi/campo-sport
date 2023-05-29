package repository

import (
	"database/sql"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/stretchr/testify/suite"
)

var dummyUser = []req.User{
	{
		Id:          1,
		Username:    "user 1",
		Password:    "password 1",
		Email:       "user1@gmail.com",
		PhoneNumber: "08145323432",
	},
}

var dummyUserRes = []res.GetAllUser{
	{
		Id:          1,
		Username:    "user 1",
		Email:       "user1@gmail.com",
		PhoneNumber: "08145323432",
		IsVerified:  true,
	},
}

type UserRepositoryTestSuite struct {
	suite.Suite
	dbMock  *sql.DB
	sqlMock sqlmock.Sqlmock
}

func (suite *UserRepositoryTestSuite) SetupTest() {
	dbMock, sqlMock, err := sqlmock.New()
	if err != nil {
		suite.T().Fatalf("an error '%s' was not expected when opening a stub db connection", err)
	}
	suite.dbMock = dbMock
	suite.sqlMock = sqlMock
}

func (suite *UserRepositoryTestSuite) TearDownTest() {
	err := suite.dbMock.Close()
	if err != nil {
		return
	}
}

func (suite *UserRepositoryTestSuite) TestUserRepository_InsertVoucher() {
	rows := sqlmock.NewRows([]string{"id"})
	rows.AddRow(123)
	//suite.sqlMock.ExpectPrepare("INSERT into courts").ExpectQuery().WithArgs(dummyCourt[0].CourtName, dummyCourt[0].Description, dummyCourt[0].IsAvailable, dummyCourt[0].CourtPrice).WillReturnRows(rows)
	suite.sqlMock.ExpectPrepare("INSERT INTO users \\(username, phone_number, password, email\\) VALUES \\(\\$1, \\$2, \\$3, \\$4 \\) RETURNING id").
		ExpectQuery().
		WithArgs(dummyUser[0].Username, dummyUser[0].PhoneNumber, dummyUser[0].Password, dummyUser[0].Email).
		WillReturnRows(rows)

	repo := NewUserRepository(suite.dbMock)

	u, err := repo.InsertUser(&dummyUser[0])
	suite.Nil(err)
	suite.Equal(int(123), u)
}

func (suite *UserRepositoryTestSuite) TestUserRepository_FindAllVoucher() {
	rows := sqlmock.NewRows([]string{"id", "username", "phone_number", "email", "is_verified"})
	for _, d := range dummyUserRes {
		rows.AddRow(d.Id, d.Username, d.PhoneNumber, d.Email, d.IsVerified)
	}
	suite.sqlMock.ExpectQuery("SELECT id, username, phone_number, email, is_verified  FROM users").
		WillReturnRows(rows)

	repo := NewUserRepository(suite.dbMock)
	users, err := repo.GetAllUsers()
	suite.Nil(err)
	suite.Equal(dummyUserRes, users)
}

// func (suite *UserRepositoryTestSuite) TestUserRepository_FindByEmail() {
// 	rows := sqlmock.NewRows([]string{"id", "username", "phone_number", "email", "is_verified"})
// 	for _, d := range dummyUserRes {
// 		rows.AddRow(d.Id, d.Username, d.PhoneNumber, d.Email, d.IsVerified)
// 	}
// 	suite.sqlMock.ExpectQuery("SELECT id, username, phone_number, email, is_verified FROM users").
// 		WillReturnRows(rows)

// 	repo := NewUserRepository(suite.dbMock)
// 	users, err := repo.FindUserByEmail(dummyUser) // Changed method name to FindByEmail
// 	suite.Require().NoError(err)     // Changed suite.Nil(err) to suite.Require().NoError(err)
// 	suite.Equal(dummyUserRes, users)
// }

func TestCourtRepositorySuite(t *testing.T) {
	suite.Run(t, new(CourtRepositoryTestSuite))
}
