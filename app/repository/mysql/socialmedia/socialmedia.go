package socialmedia

import "mygram/app/models/mysql"

type RepositorySocialMedia interface {
	Create(in mysql.SocialMedia) (out mysql.SocialMedia, err error)
}
