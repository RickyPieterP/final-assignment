package comment

import (
	"mygram/app/adapter/database"
	"mygram/app/models/mysql"
)

type commentRepo struct {
	mysql *database.MySQL
}

func NewCommentRepo(mysql *database.MySQL) RepositoryComment {
	return &commentRepo{
		mysql: mysql,
	}
}

func (c commentRepo) Find(in int) {

}

func (c commentRepo) Create(in mysql.Comment) {

}

func (c commentRepo) Update(in mysql.Comment) {

}

func (c commentRepo) Delete(in int) {

}