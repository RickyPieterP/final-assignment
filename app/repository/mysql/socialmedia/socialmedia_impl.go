package socialmedia

import "mygram/app/adapter/database"

type socialmediaRepo struct {
	mysql *database.MySQL
}

func NewRepositorySocialMedia(mysql *database.MySQL) RepositorySocialMedia {
	return &socialmediaRepo{
		mysql: mysql,
	}
}
