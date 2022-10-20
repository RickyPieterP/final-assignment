package photo

import "mygram/app/models/mysql"

type RepositoryPhoto interface {
	Create(in *mysql.Photo) (out *mysql.Photo, err error)
	Find(in int) (out []mysql.Photo, err error)
	FindOne(in int) (out *mysql.Photo, err error)
	Update(in *mysql.Photo) (out *mysql.Photo, err error)
	Delete(in int) (out bool, err error)
}
