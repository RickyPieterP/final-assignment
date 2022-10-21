package comment

import "mygram/app/models/mysql"

type RepositoryComment interface {
	Create(in *mysql.Comment) (out *mysql.Comment, err error)
	Find(in int) (out []mysql.Comment, err error)
	FindOne(in int) (out *mysql.Comment, err error)
	Update(in *mysql.Comment) (out *mysql.Comment, err error)
	Delete(in int) (out bool, err error)
}
