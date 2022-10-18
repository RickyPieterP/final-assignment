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
	err = u.mysql.Debug().Save(&in).Error
	out = in
	return
}

func (u *userRepo) Find(in mysql.User) (out mysql.User, err error) {
	err = u.mysql.Where("email = ?", in.Email).First(&out).Error
	return
}

func (*userRepo) DeleteUser() {}
