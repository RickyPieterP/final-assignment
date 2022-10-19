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

func (p *photoRepo) Find(in int) {

}

func (p *photoRepo) Create(in *mysql.Photo) (out *mysql.Photo, err error) {
	err = p.mysql.Create(in).Error
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (p *photoRepo) Update(in mysql.Photo) {
	
}

func (p *photoRepo) Delete(in int) {

}