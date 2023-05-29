package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/usecase"
	"github.com/satriabagusi/campo-sport/pkg/helper"
	"github.com/satriabagusi/campo-sport/pkg/token"
)

type CourtHandler interface {
	InsertCourt(*gin.Context)
	UpdateCourt(*gin.Context)
	DeleteCourt(*gin.Context)
	FindCourtByID(*gin.Context)
	FindCourtByCourtName(*gin.Context)
	GetAllCourts(*gin.Context)
}

type courtHandler struct {
	courtUsecase usecase.CourtUsecase
}

func NewCourtHandler(courtUsecase usecase.CourtUsecase) CourtHandler {
	return &courtHandler{courtUsecase}
}

func (h *courtHandler) InsertCourt(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorized", nil)
		return
	}
	var newCourt entity.Court
	if err := c.ShouldBindJSON(&newCourt); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request,data required!", nil)
		return
	}

	courtInDb, _ := h.courtUsecase.FindCourtByCourt(newCourt.CourtName)
	if courtInDb != nil {
		helper.Response(c, http.StatusConflict, "Court already exits!", nil)
		return
	}

	result, err := h.courtUsecase.InsertCourt(&newCourt)
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

	helper.Response(c, http.StatusOK, "Success", result)
}

func (h *courtHandler) UpdateCourt(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorizes", nil)
		return
	}
	var updateCourt entity.Court
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	updateCourt.Id = idInt

	userInDb, _ := h.courtUsecase.FindCourtById(idInt)
	if userInDb == nil {
		helper.Response(c, http.StatusNotFound, "Court not found!", nil)
		return
	}

	if err := c.ShouldBindJSON(&updateCourt); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request! data required", nil)
		return
	}
	_, err := h.courtUsecase.UpdateCourt(&updateCourt)
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

	helper.Response(c, http.StatusOK, "Court successfully Updated!", nil)

}

func (h *courtHandler) DeleteCourt(c *gin.Context) {
	user := c.MustGet("userinfo").(*token.MyCustomClaims)

	if user.UserRole != 1 {
		helper.Response(c, http.StatusUnauthorized, "Unauthorizes", nil)
		return
	}
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request! data required!", nil)
		return
	}

	_, err = h.courtUsecase.FindCourtById(id)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "Court not found!", nil)
		return
	}

	err = h.courtUsecase.DeleteCourt(&entity.Court{Id: id})
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "Court has been deleted", nil)

}

func (h *courtHandler) FindCourtByID(c *gin.Context) {

	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad request!", nil)
		return
	}

	result, err := h.courtUsecase.FindCourtById(id)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "Court not found!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)

}

func (h *courtHandler) FindCourtByCourtName(c *gin.Context) {

	voucherCode := c.Query("court_name")

	result, err := h.courtUsecase.FindCourtByCourt(voucherCode)
	if err != nil {
		helper.Response(c, http.StatusNotFound, "Court not found", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)

}

func (h *courtHandler) GetAllCourts(c *gin.Context) {
	result, err := h.courtUsecase.GetAllCourts()
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error!", nil)
		return
	}

	helper.Response(c, http.StatusOK, "OK", result)
}
