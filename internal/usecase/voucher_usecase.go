package usecase

import (
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/repository"
)

type VoucherUsecase interface {
	GetAllVoucher() ([]entity.Voucher, error)
}

type voucherUsecase struct {
	voucherRepository repository.VoucherRepository
}

func NewVoucherUsecase(voucherRepository repository.VoucherRepository) VoucherUsecase {
	return &voucherUsecase{voucherRepository}
}

func (u *voucherUsecase) GetAllVoucher() ([]entity.Voucher, error) {
	return u.voucherRepository.GetAllVoucher()
}
