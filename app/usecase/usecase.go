package usecase

import (
	"mygram/app/usecase/request"
	"mygram/app/usecase/response"
)

type Usecase interface {
	RegisterUser(in request.RegisterUserReq) (out response.RegisterUserRes, httpStatus int)
}
