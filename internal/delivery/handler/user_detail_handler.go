package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/usecase"
)

type UserDetailHandler interface {
	UploadCredential(*gin.Context)
	GetAllUserDetail(*gin.Context)
}

type userDetailHandler struct {
	userDetailUsecase usecase.UserDetailUsecase
	userUsecase       usecase.UserUsecase
}

func NewUserDetailHandler(userDetailUsecase usecase.UserDetailUsecase, userUsecase usecase.UserUsecase) UserDetailHandler {
	return &userDetailHandler{userDetailUsecase, userUsecase}
}

func (h *userDetailHandler) UploadCredential(c *gin.Context) {

	var user req.UserProfile
	id := c.Query("id")
	idInt, _ := strconv.Atoi(id)
	user.UserId = idInt

	// userInDb, _ := h.userUsecase.FindUserById(user.UserId)

	// if userInDb == nil {
	// 	c.JSON(http.StatusBadRequest, gin.H{"error": "user belum terdaftar"})
	// 	return
	// }

	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user.File, _ = c.FormFile("file")
	//file, _ = fileHeader.Open()

	updatedProfile, err := h.userDetailUsecase.UploadCredential(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	webResponse := res.WebResponse{
		Code:   201,
		Status: "Created",
		Data:   updatedProfile,
	}
	c.JSON(http.StatusOK, webResponse)
}

func (h *userDetailHandler) GetAllUserDetail(c *gin.Context) {

}

// fileHeader, _ := c.FormFile("file")
// 	file, _ := fileHeader.Open()

// 	ctx := context.Background()

// 	couldService, _ := cloudinary.NewFromURL(utility.GetEnv("CLOUDINARY_URL"))

// 	result, _ := couldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: utility.GetEnv("CLOUDINARY_UPLOAD_FOLDER")})

// 	userProfileRes := res.UserProfile{
// 		Name: user.Name,
// 		Url:  result.SecureURL,
// 	}
// 	webResponse := res.WebResponse{
// 		Code:   201,
// 		Status: "Created",
// 		Data:   tempFile,
// 	}
// 	c.JSON(http.StatusOK, webResponse)
