package socialmedia

import (
	"mygram/app/adapter/database"
	"mygram/app/models/mysql"
)

type socialmediaRepo struct {
	mysql *database.MySQL
}

func NewRepositorySocialMedia(mysql *database.MySQL) RepositorySocialMedia {
	return &socialmediaRepo{
		mysql: mysql,
	}
}

func (s socialmediaRepo) Find(id int) {
	
}

func (s socialmediaRepo) Create(in mysql.SocialMedia) {

}

func (s socialmediaRepo) Update(in mysql.SocialMedia) {

}

func (s socialmediaRepo) Delete(id int) {

}