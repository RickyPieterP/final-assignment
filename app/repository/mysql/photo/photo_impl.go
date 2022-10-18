package photo

import (
	"mygram/app/adapter/database"
	"mygram/app/models/mysql"
)

type photoRepo struct {
	mysql *database.MySQL
}

func NewRepositoryPhoto(mysql *database.MySQL) RepositoryPhoto {
	return &photoRepo{
		mysql: mysql,
	}
}

func (p photoRepo) Find(in int) {

}

func (p photoRepo) Create(in mysql.Photo) {

}

func (p photoRepo) Update(in mysql.Photo) {

}

func (p photoRepo) Delete(in int) {

}