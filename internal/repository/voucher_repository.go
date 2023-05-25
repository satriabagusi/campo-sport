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
func (r *voucherRepository) InsertVoucher(voucher *entity.Voucher) (*entity.Voucher, error) {
	stmt, err := r.db.Prepare("INSERT INTO vouchers (voucher_code, is_available, discount) VALUES ($1, $2, $3) RETURNING id")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	var voucherID int
	err = stmt.QueryRow(voucher.VoucherCode, voucher.IsAvailable, voucher.Discount).Scan(&voucherID)
	if err != nil {
		return nil, err
	}

	voucher.Id = voucherID
	return voucher, nil
}
func (r *voucherRepository) UpdateVoucher(voucher *entity.Voucher) (*entity.Voucher, error) {

	stmt, err := r.db.Prepare("UPDATE vouchers SET voucher_code = $1, is_available = $2, discount =$3 WHERE id = $4")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(voucher.VoucherCode, voucher.IsAvailable, voucher.Discount, voucher.Id)
	if err != nil {
		return nil, err
	}
	return voucher, nil

}

func (r *voucherRepository) DeleteVoucher(voucher *entity.Voucher) error {
	stmt, err := r.db.Prepare("DELETE FROM vouchers WHERE id = $1")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(voucher.Id)
	if err != nil {
		return err
	}

	return nil
}
func (r *voucherRepository) FindVoucherById(id int) (*entity.Voucher, error) {
	var voucher entity.Voucher
	stmtm, err := r.db.Prepare("SELECT voucher_code, is_available, discount FROM vouchers WHERE id = $1")
	if err != nil {
		return nil, err
	}

	defer stmtm.Close()
	row := stmtm.QueryRow(id)
	err = row.Scan(&voucher.VoucherCode, &voucher.IsAvailable, &voucher.Discount)
	if err != nil {
		return nil, err
	}
	return &voucher, nil
}
func (r *voucherRepository) FindUserByVoucher(voucherCode string) (*entity.Voucher, error) {
	var voucher entity.Voucher
	stmt, err := r.db.Prepare("SELECT id, voucher_code, is_available, discount FROM vouchers WHERE  voucher_code = $1")
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	row := stmt.QueryRow(voucherCode)
	err = row.Scan(&voucher.Id, &voucher.VoucherCode, &voucher.IsAvailable, &voucher.Discount)
	if err != nil {
		return nil, err
	}

	return &voucher, nil

}
func (r *voucherRepository) GetAllVoucher() ([]entity.Voucher, error) {
	var vouchers []entity.Voucher
	rows, err := r.db.Query("SELECT id, voucher_code, is_available, discount FROM vouchers")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	for rows.Next() {
		var voucher entity.Voucher
		err := rows.Scan(&voucher.Id, &voucher.VoucherCode, &voucher.IsAvailable, &voucher.Discount)
		if err != nil {
			return nil, err
		}
		vouchers = append(vouchers, voucher)
	}

	return vouchers, nil
}
