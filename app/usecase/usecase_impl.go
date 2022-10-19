package usecase

import (
	"fmt"
	"mygram/app/models/mysql"
	"mygram/app/repository/mysql/comment"
	"mygram/app/repository/mysql/photo"
	"mygram/app/repository/mysql/socialmedia"
	"mygram/app/repository/mysql/user"
	"mygram/app/usecase/request"
	"mygram/app/usecase/response"
	"net/http"
	"strconv"
	"time"
)

type usecase struct {
	userRepo        user.RepositoryUser
	commentRepo     comment.RepositoryComment
	photoRepo       photo.RepositoryPhoto
	socialmediaRepo socialmedia.RepositorySocialMedia
}

func NewUsecase(
	userRepo user.RepositoryUser,
	commentRepo comment.RepositoryComment,
	photoRepo photo.RepositoryPhoto,
	socialmediaRepo socialmedia.RepositorySocialMedia,
) Usecase {
	return &usecase{
		userRepo:        userRepo,
		commentRepo:     commentRepo,
		photoRepo:       photoRepo,
		socialmediaRepo: socialmediaRepo,
	}
}

func (u *usecase) RegisterUser(in request.RegisterUserReq) (out response.RegisterUserRes, httpStatus int) {
	var sqlUser mysql.AddUser

	password, _ := GeneratePassword([]byte(in.Password))

	sqlUser.Username = in.Username
	sqlUser.Email = in.Email
	sqlUser.Password = password
	sqlUser.Age = in.Age
	sqlUser.Updated_At = time.Now()

	userData, err := u.userRepo.SaveOrUpdate(sqlUser)
	if err != nil {
		httpStatus = 500
		return
	}

	out.Age = userData.Age
	out.Email = userData.Email
	// out.ID = userData.Id
	out.Username = userData.Username
	httpStatus = 201

	return
}

func (u *usecase) LoginUser(in request.LoginUserReq) (out *response.LoginUserResp, httpStatus int, err error) {
	var sqlUser mysql.User
	sqlUser.Email = in.Email
	user, err := u.userRepo.Find(sqlUser)

	if err != nil {
		return nil, 400, err
	}

	err = ValidatePassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		err = fmt.Errorf("%s", "your not allowed")
		return nil, http.StatusNotAcceptable, err
	}

	id := strconv.Itoa(user.Id)

	jwtToken := request.JWTToken{
		Id:       id,
		Email:    user.Email,
		Username: user.Username,
	}

	token, err := GenerateToken(jwtToken)
	if err != nil {
		return nil, 500, err
	}
	res := &response.LoginUserResp{
		Token: token,
	}

	return res, 200, nil
}

func (u *usecase) UpdateUser(in request.UpdateUserReq) {

}
