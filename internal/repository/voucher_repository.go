package repository

import (
	"database/sql"

	"github.com/satriabagusi/campo-sport/internal/entity"
)

type VoucherRepository interface {
	GetAllVoucher() ([]entity.Voucher, error)
}

type voucherRepository struct {
	db *sql.DB
}

func NewVoucherRepository(db *sql.DB) VoucherRepository {
	return &voucherRepository{db}
}

func (r *voucherRepository) GetAllVoucher() ([]entity.Voucher, error) {
	var voucher []entity.Voucher

	return voucher, nil
}
