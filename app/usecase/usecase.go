package usecase

import (
	"mygram/app/usecase/request"
	"mygram/app/usecase/response"
)

type Usecase interface {
	RegisterUser(in request.RegisterUserReq) (out response.RegisterUserRes, httpStatus int)
	LoginUser(in request.LoginUserReq) (out *response.LoginUserResp, httpStatus int, err error)
	CreatePhoto(in *request.CreatePhotoReq) (out *response.CreatePhotoResp, httpStatus int, err error)
}
