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
		return
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
	user, err := u.userRepo.Find(sqlUser)

	if err != nil {
		return nil, 400, err
	}

	err = ValidatePassword([]byte(user.Password), []byte(in.Password))
	if err != nil {
		err = fmt.Errorf("%s", "your not allowed")
		return nil, http.StatusNotAcceptable, err
	}

	// id := strconv.Itoa(user.Id)

	jwtToken := request.JWTToken{
		Id:       user.Id,
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
	// var sqlUser mysql.User
}

func (u *usecase) CreatePhoto(in *request.CreatePhotoReq) (out *response.CreatePhotoResp, httpStatus int, err error) {
	fmt.Println(in.UserId, "in user id")
	photo := &mysql.Photo{
		UserId:   in.UserId,
		Title:    in.Title,
		Caption:  in.Caption,
		PhotoUrl: in.PhotoUrl,
	}

	res, err := u.photoRepo.Create(photo)
	if err != nil {
		fmt.Println(err, "error di create photo")
		return nil, 400, err
	}

	resp := &response.CreatePhotoResp{
		Id:        res.Id,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		CreatedAt: res.Created_Date,
		UserId:    in.UserId,
	}
	return resp, 200, nil
}

func (u *usecase) FindPhoto(in *request.FindReq) (out []response.FindPhotoResp, httpStatus int, err error) {
	user := mysql.User{
		Id: in.UserId,
	}
	user, err = u.userRepo.FindById(user)
	if err != nil {
		fmt.Println(err, "error di find")
		return
	}
	singleUser := response.UserPhoto{
		Email:    user.Email,
		Username: user.Username,
	}
	fmt.Println(singleUser, "single user")
	res, err := u.photoRepo.Find(in.UserId)
	for i := 0; i < len(res); i++ {

		single := response.FindPhotoResp{
			Id:        res[i].Id,
			Title:     res[i].Title,
			Caption:   res[i].Caption,
			PhotoUrl:  res[i].PhotoUrl,
			UserId:    in.UserId,
			CreatedAt: res[i].Created_Date,
			UpdatedAt: res[i].Updated_At,
			User:      singleUser,
		}

		out = append(out, single)
	}
	return
}

func (u *usecase) UpdatePhoto(in *request.UpdatePhoto) (out *response.UpdatePhotoResp, err error) {
	photo := &mysql.Photo{
		Id:       in.Id,
		Title:    in.Title,
		Caption:  in.Caption,
		PhotoUrl: in.PhotoUrl,
	}
	check, err := u.photoRepo.FindOne(photo.Id)
	if err != nil {
		return
	}
	fmt.Println(check.UserId, "check user id")
	fmt.Println(in.UserId, "in user id")
	if check.UserId != in.UserId {
		err = fmt.Errorf("%s", "your unauthorized")
		return
	}

	res, err := u.photoRepo.Update(photo)
	fmt.Println(err, "error")
	if err != nil {
		return
	}
	fmt.Println(res, "res")
	out = &response.UpdatePhotoResp{
		Id:        res.Id,
		Title:     res.Title,
		Caption:   res.Caption,
		PhotoUrl:  res.PhotoUrl,
		UserId:    in.UserId,
		UpdatedAt: res.Updated_At,
	}
	return
}

func (u *usecase) DeletePhoto(in, user_id int) (out *response.DeletePhoto, err error) {
	check, err := u.photoRepo.FindOne(in)
	if err != nil {
		return
	}

	if check.UserId != user_id {
		err = fmt.Errorf("%s", "your unauthorized")
		return
	}

	res, err := u.photoRepo.Delete(in)
	if err != nil {
		return
	}
	if !res {
		return
	} else {
		out = &response.DeletePhoto{
			Message: "Your photo has been successfully deleted",
		}
		return
	}
}

func (u *usecase) CreateComment(in *request.CreateCommentReq) (out *response.CreateCommentResp, httpStatus int, err error) {
	fmt.Println(in.Message, "in user id")
	comment := &mysql.Comment{
		UserId:  in.UserId,
		PhotoId: in.PhotoId,
		Message: in.Message,
	}

	res, err := u.commentRepo.Create(comment)
	if err != nil {
		fmt.Println(err, "error di create photo")
		return nil, 400, err
	}
	fmt.Println(res, "res")
	resp := &response.CreateCommentResp{
		Id:        res.Id,
		Mesage:    res.Message,
		CreatedAt: res.Created_Date,
		UserId:    in.UserId,
		PhotoId:   res.PhotoId,
	}
	return resp, 200, nil
}

func (u *usecase) FindComment(in *request.FindCommentReq) (out []response.FindCommentResp, httpStatus int, err error) {
	user := mysql.User{
		Id: in.UserId,
	}
	user, err = u.userRepo.FindById(user)
	if err != nil {
		fmt.Println(err, "error di find")
		return
	}
	singleUser := response.UserComment{
		Id:       user.Id,
		Email:    user.Email,
		Username: user.Username,
	}
	fmt.Println(singleUser, "single user")
	res, err := u.commentRepo.Find(in.UserId)
	for i := 0; i < len(res); i++ {
		photo, err := u.photoRepo.FindOne(res[i].PhotoId)
		if err != nil {
			fmt.Println(err, "error")
		}
		photoResp := response.PhotoComment{
			Id:      photo.Id,
			Title:   photo.Title,
			Caption: photo.Caption,
			PhotoUrl: photo.PhotoUrl,
			UserId:  in.UserId,
		}
		single := response.FindCommentResp{
			Id:        res[i].Id,
			Message:   res[i].Message,
			PhotoId:   res[i].PhotoId,
			UserId:    in.UserId,
			CreatedAt: res[i].Created_Date,
			UpdatedAt: res[i].Updated_At,
			User:      singleUser,
			Photo:     photoResp,
		}

		out = append(out, single)
	}
	return
}

func (u *usecase) UpdateComment(in *request.UpdateCommentReq) (out *response.UpdateCommentResp, err error) {
	comment := &mysql.Comment{
		Id:       in.Id,
		Message: in.Message,
		UserId: in.UserId,
	}
	check, err := u.commentRepo.FindOne(comment.Id)
	if err != nil {
		return
	}
	fmt.Println(check.UserId, "check user id")
	fmt.Println(in.UserId, "in user id")
	if check.UserId != in.UserId {
		err = fmt.Errorf("%s", "your unauthorized")
		return
	}

	res, err := u.commentRepo.Update(comment)
	fmt.Println(err, "error")
	if err != nil {
		return
	}
	fmt.Println(res, "res")
	out = &response.UpdateCommentResp{
		Id:        res.Id,
		Message: res.Message,
		UserId: res.UserId,
		PhotoId: res.PhotoId,
		UpdatedAt: res.Updated_At,
	}
	return
}

func (u *usecase) DeleteComment(in, user_id int) (out *response.DeleteCommentResp, err error) {
	check, err := u.commentRepo.FindOne(in)
	if err != nil {
		return
	}

	if check.UserId != user_id {
		err = fmt.Errorf("%s", "your unauthorized")
		return
	}

	res, err := u.commentRepo.Delete(in)
	if err != nil {
		return
	}
	if !res {
		return
	} else {
		out = &response.DeleteCommentResp{
			Message: "Your photo has been successfully deleted",
		}
		return
	}
}

func (u *usecase) CreateSocialMedia(in request.CreateSocialMediaReq, userID any) (out response.CreateSocialMediaRes, httpStatus int, err error) {
	var sqlSocialMedia mysql.SocialMedia

	sqlSocialMedia.Name = in.Name
	sqlSocialMedia.SocialMediaUrl = in.SocialMediaURL

	data, err := u.socialmediaRepo.Create(sqlSocialMedia)

	out.ID = data.Id

	return
}
