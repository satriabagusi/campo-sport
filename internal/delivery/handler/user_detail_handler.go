package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/usecase"
	"github.com/satriabagusi/campo-sport/pkg/helper"
	"github.com/satriabagusi/campo-sport/pkg/token"
)

type UserDetailHandler interface {
	UploadCredential(*gin.Context)
}

type userDetailHandler struct {
	userDetailUsecase usecase.UserDetailUsecase
	userUsecase       usecase.UserUsecase
}

func NewUserDetailHandler(userDetailUsecase usecase.UserDetailUsecase, userUsecase usecase.UserUsecase) UserDetailHandler {
	return &userDetailHandler{userDetailUsecase, userUsecase}
}

func (h *userDetailHandler) UploadCredential(c *gin.Context) {

	var userReq req.UserProfile

	user := c.MustGet("userinfo").(*token.MyCustomClaims)
	userId := user.ID
	userReq.UserId = userId
	userReq.File, _ = c.FormFile("file")
	//log.Println(userReq.UserId, userReq.File)

	if err := c.ShouldBind(&userReq); err != nil {
		helper.Response(c, http.StatusBadRequest, "Bad Request!", nil)
		return
	}

	updatedProfile, err := h.userDetailUsecase.UploadCredential(&userReq)
	if err != nil {
		helper.Response(c, http.StatusInternalServerError, "Server error!", nil)
		return
	}

	helper.Response(c, http.StatusCreated, "Created", updatedProfile)
}
