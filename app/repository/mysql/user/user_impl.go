package user

import (
	"mygram/app/adapter/database"
	"mygram/app/models/mysql"
)

type userRepo struct {
	mysql *database.MySQL
}

func NewUserRepo(mysql *database.MySQL) RepositoryUser {
	return &userRepo{
		mysql: mysql,
	}
}

func (u *userRepo) SaveOrUpdate(in mysql.AddUser) (out mysql.User, err error) {
	err = u.mysql.Create(&in).Error
	return
}

func (u *userRepo) Find(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Where("email = ?", in.Email).First(&out).Error
	return
}

func (u *userRepo) FindById(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Where("id = ?", in.Id).First(&out).Error
	return
}

func (*userRepo) DeleteUser() {}
