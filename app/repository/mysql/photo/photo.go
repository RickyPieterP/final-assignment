package photo

import "mygram/app/models/mysql"

type RepositoryPhoto interface {
	Create(in *mysql.Photo) (out *mysql.Photo, err error)
}
