package comment

import "mygram/app/adapter/database"

type commentRepo struct {
	mysql *database.MySQL
}

func NewCommentRepo(mysql *database.MySQL) RepositoryComment {
	return &commentRepo{
		mysql: mysql,
	}
}
