package usecase

import (
	"context"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/satriabagusi/campo-sport/internal/entity"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/req"
	"github.com/satriabagusi/campo-sport/internal/entity/dto/res"
	"github.com/satriabagusi/campo-sport/internal/repository"
	"github.com/satriabagusi/campo-sport/pkg/utility"
)

type UserDetailUsecase interface {
	UploadCredential(*req.UserProfile) (*res.UserProfile, error)
	GetAllUserDetail() ([]entity.UserDetail, error)
}

type userDetailUsecase struct {
	userDetailRepository repository.UserDetailRepository
}

func NewUserDetailUsecase(userDetailRepository repository.UserDetailRepository) UserDetailUsecase {
	return &userDetailUsecase{userDetailRepository}
}
func (u *userDetailUsecase) UploadCredential(updatedUser *req.UserProfile) (*res.UserProfile, error) {

	file, _ := updatedUser.File.Open()
	ctx := context.Background()
	couldService, _ := cloudinary.NewFromURL(utility.GetEnv("CLOUDINARY_URL"))
	result, _ := couldService.Upload.Upload(ctx, file, uploader.UploadParams{Folder: utility.GetEnv("CLOUDINARY_UPLOAD_FOLDER")})

	userProfileRes := res.UserProfile{
		User_id: updatedUser.UserId,
		Url:  result.SecureURL,
	}
	repo, err := u.userDetailRepository.UploadCretential(&userProfileRes)
	if err != nil {
		return nil, err
	}
	return repo, nil
}

func (u *userDetailUsecase) GetAllUserDetail() ([]entity.UserDetail, error) {
	return u.userDetailRepository.GetAllUserDetail()
}
