package socialmedia

import "mygram/app/models/mysql"

type RepositorySocialMedia interface {
	SaveOrUpdate(in mysql.SocialMedia) (out mysql.SocialMedia, err error)
	Find(in int) (out mysql.SocialMedia, err error)
	FindById(in int) (out mysql.SocialMedia, err error)
	Delete(in int) (err error)
}
