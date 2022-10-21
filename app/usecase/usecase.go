package usecase

import (
	"mygram/app/usecase/request"
	"mygram/app/usecase/response"
)

type Usecase interface {
	RegisterUser(in request.RegisterUserReq) (out response.RegisterUserRes, httpStatus int)
	LoginUser(in request.LoginUserReq) (out *response.LoginUserResp, httpStatus int, err error)
	UpdateUser(in request.UpdateUserReq) (out *response.UpdateUserRes, httpStatus int, err error)
	DeleteUser(in Token) (out *response.DeleteUser, httpStatus int, err error)
	CreatePhoto(in *request.CreatePhotoReq) (out *response.CreatePhotoResp, httpStatus int, err error)
	FindPhoto(in *request.FindReq) (out []response.FindPhotoResp, httpStatus int, err error)
	UpdatePhoto(in *request.UpdatePhoto) (out *response.UpdatePhotoResp, err error)
	DeletePhoto(in, user_id int) (out *response.DeletePhoto, err error)
	CreateSocialMedia(in request.CreateSocialMediaReq, userID any) (out response.CreateSocialMediaRes, httpStatus int, err error)
	// FindSocialMedia(in request.FindSocialMediaReq) (out response.FindSocialMediaRes, err error)
}
