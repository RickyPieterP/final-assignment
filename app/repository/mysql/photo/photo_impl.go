package photo

import "mygram/app/adapter/database"

type photoRepo struct {
	mysql *database.MySQL
}

func NewRepositoryPhoto(mysql *database.MySQL) RepositoryPhoto {
	return &photoRepo{
		mysql: mysql,
	}
}
