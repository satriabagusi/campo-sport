package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type VoucherUsecase interface {
	InsertVoucher(*entity.Voucher) (*entity.Voucher, error)
	UpdateVoucher(*entity.Voucher) (*entity.Voucher, error)
	DeleteVoucher(*entity.Voucher) error
	FindVoucherById(int) (*entity.Voucher, error)
	FindVoucherByVoucher(string) (*entity.Voucher, error)
	GetAllVoucher() ([]entity.Voucher, error)
}

type voucherUsecase struct {
	voucherRepository repository.VoucherRepository
}

func NewVoucherUsecase(voucherRepository repository.VoucherRepository) VoucherUsecase {
	return &voucherUsecase{voucherRepository}
}

func (u *voucherUsecase) InsertVoucher(voucher *entity.Voucher) (*entity.Voucher, error) {
	return u.voucherRepository.InsertVoucher(voucher)
}
func (u *voucherUsecase) UpdateVoucher(voucher *entity.Voucher) (*entity.Voucher, error) {
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
