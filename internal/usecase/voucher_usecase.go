package usecase

import (
	"github.com/go-playground/validator/v10"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type VoucherUsecase interface {
	InsertVoucher(*entity.Voucher) (*entity.Voucher, error)
	UpdateVoucher(*req.UpdateVoucher) (*req.UpdateVoucher, error)
	DeleteVoucher(*entity.Voucher) error
	FindVoucherById(int) (*entity.Voucher, error)
	FindVoucherByVoucher(string) (*entity.Voucher, error)
	GetAllVoucher() ([]entity.Voucher, error)
}

type voucherUsecase struct {
	voucherRepository repository.VoucherRepository
	validate          *validator.Validate
}

func NewVoucherUsecase(voucherRepository repository.VoucherRepository, validate *validator.Validate) VoucherUsecase {
	return &voucherUsecase{voucherRepository, validate}
}

func (u *voucherUsecase) InsertVoucher(voucher *entity.Voucher) (*entity.Voucher, error) {
	err := u.validate.Struct(voucher)
	if err != nil {
		return nil, err
	}
	return u.voucherRepository.InsertVoucher(voucher)
}
func (u *voucherUsecase) UpdateVoucher(voucher *req.UpdateVoucher) (*req.UpdateVoucher, error) {
	err := u.validate.Struct(voucher)
	if err != nil {
		return nil, err
	}
	return u.voucherRepository.UpdateVoucher(voucher)
}
func (u *voucherUsecase) DeleteVoucher(user *entity.Voucher) error {
	return u.voucherRepository.DeleteVoucher(user)
}
func (u *voucherUsecase) FindVoucherById(id int) (*entity.Voucher, error) {
	return u.voucherRepository.FindVoucherById(id)
}
func (u *voucherUsecase) FindVoucherByVoucher(voucherCode string) (*entity.Voucher, error) {
	return u.voucherRepository.FindUserByVoucher(voucherCode)
}

func (u *voucherUsecase) GetAllVoucher() ([]entity.Voucher, error) {
	return u.voucherRepository.GetAllVoucher()
}
