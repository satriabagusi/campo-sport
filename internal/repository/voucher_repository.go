package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type VoucherRepository interface {
	InsertVoucher(*entity.Voucher) (*entity.Voucher, error)
	UpdateVoucher(*entity.Voucher) (*entity.Voucher, error)
	DeleteVoucher(*entity.Voucher) error
	FindVoucherById(int) (*entity.Voucher, error)
	FindUserByVoucher(string) (*entity.Voucher, error)
	GetAllVoucher() ([]entity.Voucher, error)
}

type voucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return &voucherRepository{db}
}
func (r *voucherRepository) InsertVoucher(*entity.Voucher) (*entity.Voucher, error) {
	panic("implement me")
}
func (r *voucherRepository) UpdateVoucher(*entity.Voucher) (*entity.Voucher, error) {
	panic("implement me")
}

func (r *voucherRepository) DeleteVoucher(*entity.Voucher) error {
	panic("implement me")
}
func (r *voucherRepository) FindVoucherById(int) (*entity.Voucher, error) {
	panic("implement me")
}
func (r *voucherRepository) FindUserByVoucher(string) (*entity.Voucher, error) {
	panic("implement me")
}
func (r *voucherRepository) GetAllVoucher() ([]entity.Voucher, error) {
	var voucher []entity.Voucher

	return voucher, nil
}
