package user

import "mygram/app/models/mysql"

type RepositoryUser interface {
	SaveOrUpdate(in mysql.User) (out mysql.User, err error)
	Find(in mysql.User) (out mysql.User, err error)
	FindById(in mysql.User) (out mysql.User, err error)
	UpdateUser(in mysql.User) (out mysql.User, err error)
	DeleteUser(in mysql.User) error
}
