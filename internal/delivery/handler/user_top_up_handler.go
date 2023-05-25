/*
Author: Satria Bagus(satria.bagus18@gmail.com)
user_top_up_handler.go (c) 2023
Desc: description
Created:  2023-05-24T18:10:28.456Z
Modified: !date!
*/

package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserTopUpHandler interface {
	TopUpBalance(*gin.Context)
}

type userTopUpHandler struct {
	userTopUpUsecase usecase.UserTopUpUsecase
}

func NewUserTopUpHandler(userTopUpUsecase usecase.UserTopUpUsecase) UserTopUpHandler {
	return &userTopUpHandler{userTopUpUsecase}
}

func (h *userTopUpHandler) TopUpBalance(ctx *gin.Context) {
	var userTopUp entity.UserTopUp

	if err := ctx.ShouldBindJSON(&userTopUp); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    userTopUp,
		})
		return
	}

	result, err := h.userTopUpUsecase.TopUpBalance(&userTopUp)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    userTopUp,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code":    http.StatusAccepted,
		"message": "Successfully Created Top Up Request.",
		"data":    result,
	})
}
