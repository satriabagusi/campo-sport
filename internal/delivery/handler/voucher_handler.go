package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type VoucherHandler interface {
	InsertVoucher(*gin.Context)
	UpdateVoucher(*gin.Context)
	DeleteVoucher(*gin.Context)
	FindVoucherByID(*gin.Context)
	FindVoucherByVoucherCode(*gin.Context)
	GetAllVoucher(*gin.Context)
}

type voucherHandler struct {
	voucherUsecase usecase.VoucherUsecase
}

func NewVoucherHandler(voucherUsecase usecase.VoucherUsecase) VoucherHandler {
	return &voucherHandler{voucherUsecase}
}
func (h *voucherHandler) InsertVoucher(c *gin.Context) {
	var voucher entity.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	voucherInDb, _ := h.voucherUsecase.FindVoucherByVoucher(voucher.VoucherCode)
	if voucherInDb != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "voucher already exists"})
		return
	}

	result, err := h.voucherUsecase.InsertVoucher(&voucher)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	webResponse := res.WebResponse{
		Code:   201,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusCreated, webResponse)
}

func (h *voucherHandler) UpdateVoucher(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.voucherUsecase.FindVoucherById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
		return
	}

	var updatedVoucher entity.Voucher
	if err := c.ShouldBindJSON(&updatedVoucher); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	result, err := h.voucherUsecase.UpdateVoucher(&updatedVoucher)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update voucher"})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)

}
func (h *voucherHandler) DeleteVoucher(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.voucherUsecase.FindVoucherById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
		return
	}

	err = h.voucherUsecase.DeleteVoucher(&entity.Voucher{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "voucher has been deleted",
	}

	c.JSON(http.StatusOK, webResponse)

}
func (h *voucherHandler) FindVoucherByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.voucherUsecase.FindVoucherById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "voucher not found"})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)

}
func (h *voucherHandler) FindVoucherByVoucherCode(c *gin.Context) {
	voucherCode := c.Query("voucher_code")

	result, err := h.voucherUsecase.FindVoucherByVoucher(voucherCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "user not found"})
		return
	}
	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (h *voucherHandler) GetAllVoucher(c *gin.Context) {
	result, err := h.voucherUsecase.GetAllVoucher()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)
}
