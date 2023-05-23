package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type VoucherHandler interface {
	GetAllVoucher(*gin.Context)
}

type voucherHandler struct {
	voucherUsecase usecase.VoucherUsecase
}

func NewVoucherHandler(voucherUsecase usecase.VoucherUsecase) VoucherHandler {
	return &voucherHandler{voucherUsecase}
}

func (h *voucherHandler) GetAllVoucher(c *gin.Context) {

}
