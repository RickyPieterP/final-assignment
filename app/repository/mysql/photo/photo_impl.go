package photo

import (
	"fmt"
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

func (p *photoRepo) Find(in int) (out []mysql.Photo, err error) {
	err = p.mysql.Where("user_id = ?", in).Find(&out).Error
	return
}

func (p *photoRepo) FindOne(in int) (out *mysql.Photo, err error) {
	out = &mysql.Photo{}
	err = p.mysql.Where("id = ?", &in).First(out).Error
	fmt.Println(err, "error di findOne")
	return
}

func (p *photoRepo) Create(in *mysql.Photo) (out *mysql.Photo, err error) {
	err = p.mysql.Create(in).Error
	if err != nil {
		return nil, err
	}
	return in, nil
}

func (p *photoRepo) Update(in *mysql.Photo) (out *mysql.Photo, err error) {
	out = &mysql.Photo{}
	err = p.mysql.Where("id = ?", in.Id).Updates(in).Error
	if err != nil {
		return nil, err
	}
	p.mysql.Where("id = ?", in.Id).First(out)
	return out, nil
}

func (p *photoRepo) Delete(in int) (out bool, err error) {
	row := p.mysql.Model(&mysql.Photo{}).Where("id = ?", in).Delete(&mysql.Photo{})
	if row.RowsAffected != 0 {
		out = true
		return out, nil
	} else {
		err = row.Error
		if err != nil {
			return false, err
		}
		return false, nil
	}
}