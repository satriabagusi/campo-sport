package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/usecase"
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
	var newCourt entity.Court
	if err := c.ShouldBindJSON(&newCourt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	courtInDb, _ := h.courtUsecase.FindCourtByCourt(newCourt.CourtName)
	if courtInDb != nil {
		c.JSON(http.StatusConflict, gin.H{"error": "court already exists"})
		return
	}

	result, err := h.courtUsecase.InsertCourt(&newCourt)
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

func (h *courtHandler) UpdateCourt(c *gin.Context) {
	var updateCourt entity.Court
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	updateCourt.Id = idInt

	userInDb, _ := h.courtUsecase.FindCourtById(idInt)
	if userInDb == nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Court not found"})
		return
	}

	if err := c.ShouldBindJSON(&updateCourt); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err := h.courtUsecase.UpdateCourt(&updateCourt)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update voucher"})
		return
	}

	webResponse := res.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   "Court sucessfully updated",
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *courtHandler) DeleteCourt(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err = h.courtUsecase.FindCourtById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "court not found"})
		return
	}

	err = h.courtUsecase.DeleteCourt(&entity.Court{Id: id})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   "Court has been deleted",
	}

	c.JSON(http.StatusOK, webResponse)
}

func (h *courtHandler) FindCourtByID(c *gin.Context) {
	idParam := c.Param("id")

	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	result, err := h.courtUsecase.FindCourtById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "court not found"})
		return
	}

	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}

	c.JSON(http.StatusOK, webResponse)
}
func (h *courtHandler) FindCourtByCourtName(c *gin.Context) {
	voucherCode := c.Query("court_name")

	result, err := h.courtUsecase.FindCourtByCourt(voucherCode)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "court not found"})
		return
	}
	webResponse := res.WebResponse{
		Code:   200,
		Status: "OK",
		Data:   result,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (h *courtHandler) GetAllCourts(c *gin.Context) {
	result, err := h.courtUsecase.GetAllCourts()
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
