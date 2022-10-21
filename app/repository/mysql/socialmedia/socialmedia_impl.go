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

func (s *socialmediaRepo) Find(id int) {
	_ = s.mysql.DB.Find(id).Error
	return
}

func (s *socialmediaRepo) Create(in mysql.SocialMedia) (out mysql.SocialMedia, err error) {
	err = s.mysql.DB.Create(in).Error
	out = in
	return
}

func (s *socialmediaRepo) Update(in mysql.SocialMedia) {

}

func (s *socialmediaRepo) Delete(id int) {

}
