package socialmedia

import "mygram/app/models/mysql"

type RepositorySocialMedia interface {
	SaveOrUpdate(in mysql.SocialMedia) (out mysql.SocialMedia, err error)
}
