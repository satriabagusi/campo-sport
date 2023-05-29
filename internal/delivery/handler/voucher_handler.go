package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/usecase"
	"github.com/satriabagusi/campo-sport/pkg/helper"
	"github.com/satriabagusi/campo-sport/pkg/token"
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
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var voucher entity.Voucher
	if err := c.ShouldBindJSON(&voucher); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request! Data required", nil)
		return
	}

	voucherInDb, _ := h.voucherUsecase.FindVoucherByVoucher(voucher.VoucherCode)
	if voucherInDb != nil {
		helper.Response(c, http.StatusConflict, "voucher already exist", nil)
		return
	}

	result, err := h.voucherUsecase.InsertVoucher(&voucher)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Handle validation errors
			validationErrors := make(map[string]string)

			for _, e := range validationErrs {
				validationErrors[e.Field()] = e.Tag()
			}

			helper.Response(c, http.StatusBadRequest, "Validation Error", validationErrors)
			return
		}

		helper.Response(c, http.StatusInternalServerError, "Server Error!", nil)
		return
	}

	helper.Response(c, http.StatusCreated, "OK", result)
}

func (h *voucherHandler) UpdateVoucher(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}

	var updateVoucher req.UpdateVoucher
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	updateVoucher.Id = idInt

	userInDb, _ := h.voucherUsecase.FindVoucherById(idInt)
	if userInDb == nil {
		helper.Response(c, http.StatusNotFound, "voucher not found", nil)
		return
	}

	if err := c.ShouldBindJSON(&updateVoucher); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request!", nil)
		return
	}
	_, err := h.voucherUsecase.UpdateVoucher(&updateVoucher)
	if err != nil {
		if validationErrs, ok := err.(validator.ValidationErrors); ok {
			// Handle validation errors
			validationErrors := make(map[string]string)

			for _, e := range validationErrs {
				validationErrors[e.Field()] = e.Tag()
			}

			helper.Response(c, http.StatusBadRequest, "Validation Error", validationErrors)
			return
		}

		helper.Response(c, http.StatusInternalServerError, "Server Error!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "Voucher successfully updated!", nil)

}
func (h *voucherHandler) DeleteVoucher(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request!", nil)
		return
	}

	_, err = h.voucherUsecase.FindVoucherById(id)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "voucher not found", nil)
		return
	}

	err = h.voucherUsecase.DeleteVoucher(&entity.Voucher{Id: id})
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "failed to delete voucher", nil)
		return
	}

	helper.Response(c, http.StatusOK, "voucher has been deleted", nil)

}
func (h *voucherHandler) FindVoucherByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request!", nil)
		return
	}

	result, err := h.voucherUsecase.FindVoucherById(id)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "Voucher not found", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)

}
func (h *voucherHandler) FindVoucherByVoucherCode(c *gin.Context) {

	voucherCode := c.Query("voucher_code")

	result, err := h.voucherUsecase.FindVoucherByVoucher(voucherCode)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "user not found", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}

func (h *voucherHandler) GetAllVoucher(c *gin.Context) {
	result, err := h.voucherUsecase.GetAllVoucher()
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "failed to get vouchers", nil)
		return
	}
	helper.Response(c, http.StatusOK, "OK", result)
}
