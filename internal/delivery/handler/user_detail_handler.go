package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserDetailHandler interface {
	GetAllUserDetail(*gin.Context)
}

type userDetailHandler struct {
	userDetailUsecase usecase.UserDetailUsecase
}

func NewUserDetailHandler(userDetailUsecase usecase.UserDetailUsecase) UserDetailHandler {
	return &userDetailHandler{userDetailUsecase}
}

func (h *userDetailHandler) GetAllUserDetail(c *gin.Context) {

}
