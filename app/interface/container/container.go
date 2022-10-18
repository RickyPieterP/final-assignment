package container

import (
	"mygram/app/adapter/database"
	"mygram/app/repository/mysql/comment"
	"mygram/app/repository/mysql/photo"
	"mygram/app/repository/mysql/socialmedia"
	"mygram/app/repository/mysql/user"
	"mygram/app/usecase"
)

type Container struct {
	MygramUsecase usecase.Usecase
	Response      string
}

func SetupContainer() Container {
	mysql := database.SetupMySQL()

	userRepo := user.NewUserRepo(mysql)
	commentRepo := comment.NewCommentRepo(mysql)
	photoRepo := photo.NewRepositoryPhoto(mysql)
	socialRepo := socialmedia.NewRepositorySocialMedia(mysql)

	mygramUsecase := usecase.NewUsecase(userRepo, commentRepo, photoRepo, socialRepo)

	return Container{
		MygramUsecase: mygramUsecase,
		Response:      "",
	}
}
