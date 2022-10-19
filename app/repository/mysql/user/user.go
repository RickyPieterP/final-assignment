package user

import "mygram/app/models/mysql"

type RepositoryUser interface {
	SaveOrUpdate(in mysql.AddUser) (out mysql.User, err error)
	Find(in mysql.User) (out mysql.User, err error)
	DeleteUser()
}
