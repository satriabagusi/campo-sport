package usecase

import (
	"errors"
	"log"
	"os"
	"testing"

	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
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

type VoucherRepositoryMock struct {
	mock.Mock
}

func (u *VoucherRepositoryMock) InsertVoucher(voucher *entity.Voucher) (*entity.Voucher, error) {
	args := u.Called(voucher)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*entity.Voucher), nil
}

func (u *VoucherRepositoryMock) UpdateVoucher(*req.UpdateVoucher) (*req.UpdateVoucher, error) {
	// TODO implement me
	panic("implement me")
}

func (u *VoucherRepositoryMock) DeleteVoucher(*entity.Voucher) error {
	// TODO implement me
	panic("implement me")
}

func (u *VoucherRepositoryMock) FindVoucherById(int) (*entity.Voucher, error) {
	// TODO implement me
	panic("implement me")
}

func (u *VoucherRepositoryMock) FindVoucherByVoucher(string) (*entity.Voucher, error) {
	// TODO implement me
	panic("implement me")
}
func (u *VoucherRepositoryMock) FindUserByVoucher(string) (*entity.Voucher, error) {
	// TODO implement me
	panic("implement me")
}

func (u *VoucherRepositoryMock) GetAllVoucher() ([]entity.Voucher, error) {
	args := u.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]entity.Voucher), nil
}

type VoucherUsecaseTestSuite struct {
	suite.Suite
	repoMock *VoucherRepositoryMock
}

func (suite *VoucherUsecaseTestSuite) SetupTest() {
	suite.repoMock = new(VoucherRepositoryMock)
}

func (suite *VoucherUsecaseTestSuite) TestInsertVoucher() {
	voucher := dummyVoucher[0]
	suite.repoMock.On("InsertVoucher", &voucher).Return(&voucher, nil)
	voucherUsecase := NewVoucherUsecase(suite.repoMock)
	result, err := voucherUsecase.InsertVoucher(&voucher)
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), &voucher, result)
}

func (suite *VoucherUsecaseTestSuite) TestInsertVoucherError() {
	voucher := dummyVoucher[0]
	suite.repoMock.On("InsertVoucher", &voucher).Return(nil, errors.New("error"))
	voucherUsecase := NewVoucherUsecase(suite.repoMock)
	result, err := voucherUsecase.InsertVoucher(&voucher)
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func (suite *VoucherUsecaseTestSuite) TestFindAllVoucher() {
	suite.repoMock.On("GetAllVoucher").Return(dummyVoucher, nil)
	voucherUsecase := NewVoucherUsecase(suite.repoMock)
	result, err := voucherUsecase.GetAllVoucher()
	assert.NoError(suite.T(), err)
	assert.Equal(suite.T(), dummyVoucher, result)
}

func (suite *VoucherUsecaseTestSuite) TestFindAllVoucherError() {
	dir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Current working directory:", dir)

	suite.repoMock.On("GetAllVoucher").Return(nil, errors.New("error"))
	voucherUsecase := NewVoucherUsecase(suite.repoMock)
	result, err := voucherUsecase.GetAllVoucher()
	assert.Error(suite.T(), err)
	assert.Nil(suite.T(), result)
}

func TestVoucherUsecaseSuite(t *testing.T) {
	suite.Run(t, new(VoucherUsecaseTestSuite))
}
