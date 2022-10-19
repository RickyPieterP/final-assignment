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

func (u *usecase) LoginUser(in request.LoginUserReq) (out *response.LoginUserResp, httpStatus int, err error) {
	var sqlUser mysql.User
	sqlUser.Email = in.Email
	fmt.Println(in.Email, "email")
	user, err := u.userRepo.Find(sqlUser)

	if err != nil {
		return nil, 400, err
	}
	
	err = ValidatePassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		err = fmt.Errorf("%s", "your not allowed")
		return nil, 400, err
	}

	jwtToken := request.JWTToken {
		Id: sqlUser.Id,
		Username: sqlUser.Username,
	}

	token, err := GenerateToken(jwtToken)
	fmt.Println(err, "error")
	if err != nil {
		return nil, 500, err
	}
	res := &response.LoginUserResp{
		Token: token,
	}
	
	return res, 200, nil
}

func (u *usecase) UpdateUser(in request.UpdateUserReq) {
	// var sqlUser mysql.User
}


func (u *usecase) CreatePhoto(in *request.CreatePhotoReq) (out *response.CreatePhotoResp, httpStatus int, err error) {
	photo := &mysql.Photo {
		UserId: in.UserId,
		Title: in.Title,
		Caption: in.Caption,
		PhotoUrl: in.PhotoUrl,
	}

	res, err := u.photoRepo.Create(photo)
	if err != nil {
		fmt.Println(err, "error di create photo")
		return nil, 400, err
	}

	resp := &response.CreatePhotoResp{
		Id: res.Id,
		Title: res.Title,
		Caption: res.Caption,
		PhotoUrl: res.PhotoUrl,
		CreatedAt: res.Created_Date,
	}
	return resp, 200, nil
}

func (u *usecase) FindPhoto() 