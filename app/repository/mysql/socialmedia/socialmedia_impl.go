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

func (s *socialmediaRepo) Find(in int) (out mysql.SocialMedia, err error) {
	_ = s.mysql.DB.Where("user_id = ?", &in).Find(&out).Error
	return
}

func (s *socialmediaRepo) SaveOrUpdate(in mysql.SocialMedia) (out mysql.SocialMedia, err error) {
	err = s.mysql.Debug().Save(&in).Error
	out = in
	return
}

func (s *socialmediaRepo) FindById(in int) (out mysql.SocialMedia, err error) {
	_ = s.mysql.DB.Where("id = ?", &in).Find(&out).Error
	return
}

func (s *socialmediaRepo) Delete(id int) (err error) {
	err = s.mysql.Delete(&mysql.SocialMedia{}, id).Error
	return
}
