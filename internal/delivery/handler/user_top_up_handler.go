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
	"github.com/satriabagusi/campo-sport/pkg/helper"
	"github.com/satriabagusi/campo-sport/pkg/token"
)

type UserTopUpHandler interface {
	TopUpBalance(*gin.Context)
	CheckBalance(*gin.Context)
	WithdrawBalance(*gin.Context)
}

type userTopUpHandler struct {
	userTopUpUsecase usecase.UserTopUpUsecase
}

func NewUserTopUpHandler(userTopUpUsecase usecase.UserTopUpUsecase) UserTopUpHandler {
	return &userTopUpHandler{userTopUpUsecase}
}

func (h *userTopUpHandler) TopUpBalance(ctx *gin.Context) {
	user := ctx.MustGet("userinfo").(*token.MyCustomClaims)
	userId := user.ID
	var userTopUp entity.UserTopUp
	userTopUp.User.Id = userId

	if !user.IsVerified {
		helper.Response(ctx, http.StatusBadRequest, "User not verified. Please complete the verification process.", nil)
		return
	}

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

func (h *userTopUpHandler) CheckBalance(c *gin.Context) {
	bookingNumber := c.Query("order_number")

	result, err := h.userTopUpUsecase.CheckBalance(bookingNumber)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    http.StatusNotFound,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"message": "Check Balance User Successfully",
		"data":    result,
	})
}

func (h *userTopUpHandler) WithdrawBalance(ctx *gin.Context) {

	user := ctx.MustGet("userinfo").(*token.MyCustomClaims)
	userId := user.ID
	var userWithdraw entity.UserWithdraw
	userWithdraw.User.Id = userId
		if !user.IsVerified {
			helper.Response(ctx, http.StatusBadRequest, "User not verified. Please complete the verification first", nil)
			return
		}

	if err := ctx.ShouldBindJSON(&userWithdraw); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"message": err.Error(),
			"data":    userWithdraw,
		})
		return
	}

	result, err := h.userTopUpUsecase.WithdrawBalance(&userWithdraw)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"code":    http.StatusInternalServerError,
			"message": err.Error(),
			"data":    userWithdraw,
		})
		return
	}

	ctx.JSON(http.StatusAccepted, gin.H{
		"code":    http.StatusAccepted,
		"message": "Successfully Created Top Up Request.",
		"data":    result,
	})
}
