package usecase

import (
	"mygram/app/models/mysql"
	"mygram/app/repository/mysql/comment"
	"mygram/app/repository/mysql/photo"
	"mygram/app/repository/mysql/socialmedia"
	"mygram/app/repository/mysql/user"
	"mygram/app/usecase/request"
	"mygram/app/usecase/response"
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
	var sqlUser mysql.User

	password, _ := GeneratePassword([]byte(in.Password))

	sqlUser.Username = in.Username
	sqlUser.Email = in.Email
	sqlUser.Password = password
	sqlUser.Age = in.Age
	sqlUser.Updated_At = time.Now()

	userData, err := u.userRepo.SaveOrUpdate(sqlUser)
	if err != nil {
		httpStatus = 500
	}

	out.Age = userData.Age
	out.Email = userData.Email
	out.ID = userData.Id
	out.Username = userData.Username
	httpStatus = 201

	return
}
