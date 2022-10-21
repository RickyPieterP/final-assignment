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

func (u *userRepo) SaveOrUpdate(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Create(&in).Error
	if err != nil {
		return in, err
	}
	return in, nil
}

func (u *userRepo) Find(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Where("email = ?", in.Email).First(&out).Error
	return
}

func (u *userRepo) FindById(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Where("id = ?", in.Id).First(&out).Error
	return
}

func (u *userRepo) UpdateUser(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Where("id = ?", in.Id).Updates(&in).Error
	out = in
	return
}

func (u *userRepo) DeleteUser(in mysql.User) (err error) {
	err = u.mysql.Where("id = ?", in.Id).Delete(&in).Error
	return
}
